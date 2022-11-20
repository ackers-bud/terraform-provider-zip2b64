package client

import (
    "archive/zip"
    "bytes"
    "encoding/base64"
    "errors"
    "io/ioutil"
)

// ZipExtract takes a Base64 string which is a representation of a zip archive.
//    the filenameToExtract - is a file in Zip Archive that needs to be
//    extracted
//
func ZipExtract(base64Encodedfile string, filenameToExtract string) (string, error) {

    // Convert the Base64 string to a stream of Bytes
    bytesInFile, err := base64.StdEncoding.DecodeString(base64Encodedfile)
    if err != nil {
        return "", err
    }

    // Pass the Bytes in the File to the Zip Decoder.
    ZipReader, err := zip.NewReader(bytes.NewReader(bytesInFile), int64(len(bytesInFile)))
    if err != nil {
        return "", err
    }

    // Go through all files in the Archive
    for _, Singlefile := range ZipReader.File {

        // fmt.Printf("Name of the File in the Archive is %s", Singlefile.Name)

        if Singlefile.Name == filenameToExtract {
            SinglefileContents, err := Singlefile.Open()

            if err != nil {
                return "", err
            }

            SinglefileContentsBytes, err := ioutil.ReadAll(SinglefileContents)
            if err != nil {
                return "", err
            }
            _ = SinglefileContents.Close()

            // fmt.Printf("File Contents are:\n%s\n\n", SinglefileContentsBytes)
            return base64.StdEncoding.EncodeToString(SinglefileContentsBytes), nil
        }
    }

    return "", errors.New("file not found in archive")
}
