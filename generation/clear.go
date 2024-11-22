package generation

import "os"

func clearBuildDir() error {
	entries, err := os.ReadDir(BUILD_DIR)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		err = os.RemoveAll(BUILD_DIR + "/" + entry.Name())
		if err != nil {
			return err
		}
	}
	return nil
}
