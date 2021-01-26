package node_keeper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"yu/config"
	"yu/storage/kv"
	"yu/utils/compress"
)

const CompressedFileType = ".zip"

type NodeKeeper struct {
	repoDB kv.KV
	dir    string
	port   string
}

func NewNodeKeeper(cfg *config.NodeKeeperConf) (*NodeKeeper, error) {
	dbpath := cfg.RepoDbPath
	dir := cfg.Dir
	if filepath.Dir(dbpath) == dir {
		dbpath = filepath.Join(dir, filepath.Base(dbpath))
	}
	repoDB, err := kv.NewBolt(dbpath)
	if err != nil {
		return nil, err
	}
	return &NodeKeeper{
		repoDB: repoDB,
		dir:    dir,
		port:   ":" + cfg.ServesPort,
	}, nil
}

func (n *NodeKeeper) handleFromMaster() {
	r := gin.Default()

	r.POST("/upgrade", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file error: %s", err.Error()))
			return
		}

		fname := file.Filename
		if !strings.HasSuffix(fname, CompressedFileType) {
			c.String(
				http.StatusBadRequest,
				fmt.Sprintf("the type of file(%s) is wrong", fname),
			)
			return
		}
		zipFileName := filepath.Join(n.dir, fname)
		err = c.SaveUploadedFile(file, zipFileName)
		if err != nil {
			c.String(
				http.StatusInternalServerError,
				fmt.Sprintf("save file(%s) error: %s", fname, err.Error()),
			)
			return
		}
		err = n.convertToRepo(zipFileName, fname)
		if err != nil {
			c.String(
				http.StatusBadRequest,
				fmt.Sprintf("convert file(%s) to repo error: %s", fname, err.Error()),
			)
		}

		c.String(http.StatusOK, "upload file succeed")
	})

	r.Run(n.port)
}

// zipFileName just like: path/to/yu-repo_3.zip
// 3 is the version of repo, 'yu-repo' is the name of repo
func (n *NodeKeeper) convertToRepo(zipFilePath, fname string) error {

	// repoVDir: path/to/yu-repo_3
	repoVDir := strings.TrimSuffix(zipFilePath, CompressedFileType)

	arr := strings.Split(repoVDir, "_")
	repoVersionStr := arr[len(arr)-1]
	repoVersion, err := strconv.Atoi(repoVersionStr)
	if err != nil {
		return err
	}

	// repoName: yu-repo
	repoName := strings.TrimSuffix(fname, "_"+repoVersionStr+CompressedFileType)

	// repoDir: path/to/yu-repo/3
	repoDir := filepath.Join(n.dir, repoName, repoVersionStr)
	err = os.MkdirAll(repoDir, os.ModePerm)
	if err != nil {
		return err
	}

	files, err := compress.UnzipFile(zipFilePath, repoDir)
	if err != nil {
		return err
	}

	repo, err := NewRepo(repoName, files, repoDir, repoVersion)
	if err != nil {
		return err
	}
	err = n.setRepo(repoName, repo)
	if err != nil {
		return err
	}
	return os.Remove(zipFilePath)
}

func (n *NodeKeeper) getRepo(repoName string) (*Repo, error) {
	repoByt, err := n.repoDB.Get([]byte(repoName))
	if err != nil {
		return nil, err
	}
	return decodeRepo(repoByt)
}

func (n *NodeKeeper) setRepo(repoName string, repo *Repo) error {
	repoByt, err := repo.encode()
	if err != nil {
		return err
	}
	return n.repoDB.Set([]byte(repoName), repoByt)
}
