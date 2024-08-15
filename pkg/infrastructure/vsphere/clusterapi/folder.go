package clusterapi

import (
	"context"
	"path"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/session"
)

func folderExists(ctx context.Context, dir string, session *session.Session) (*object.Folder, error) {
	/* scenarios:
	 * 1. folder exists and returns
	 * 2. folder does not exist and err and folder are nil
	 * 3. finder.Folder fails and returns folder nil and error
	 */
	var notFoundError *find.NotFoundError
	folder, err := session.Finder.Folder(ctx, dir)

	// todo: jcallen: what if the folder is not found because its the dc vm folder?

	/*
			DEBUG folder not found: /nested8-datacenter/vm
		DEBUG folder found: /dcfolder/nested8-datacenter/vm

	*/

	// scenario two
	if folder == nil && errors.As(err, &notFoundError) {
		logrus.Debugf("folder not found: %s\n", dir)
		return nil, nil
	}
	// scenario three
	if err != nil {
		return nil, err
	}
	// scenario one
	logrus.Debugf("folder found: %s\n", dir)
	return folder, nil
}

func createFolder(ctx context.Context, fullpath string, session *session.Session, vmFolder *object.Folder) (*object.Folder, error) {
	var folder *object.Folder
	var err error

	dir := path.Dir(fullpath)
	base := path.Base(fullpath)

	/*
		if base == "vm" {
			return vmFolder, nil
		}

	*/

	// if folder is nil the fullpath does not exist
	if folder, err = folderExists(ctx, dir, session); err == nil && folder == nil {
		folder, err = createFolder(ctx, dir, session, vmFolder)
		if err != nil {
			return nil, err
		}
	}

	if folder != nil && err == nil {
		logrus.Debugf("create folder: %s\n", base)
		return folder.CreateFolder(ctx, base)
	}
	return folder, err
}
