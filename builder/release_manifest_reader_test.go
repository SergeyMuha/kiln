package builder_test

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"time"

	"github.com/pivotal-cf/kiln/builder"
	"github.com/pivotal-cf/kiln/builder/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReleaseManifestReader", func() {
	var (
		filesystem *fakes.Filesystem
		reader     builder.ReleaseManifestReader
	)

	BeforeEach(func() {
		filesystem = &fakes.Filesystem{}
		reader = builder.NewReleaseManifestReader(filesystem)

		tarball := NewBuffer(bytes.NewBuffer([]byte{}))
		gw := gzip.NewWriter(tarball)
		tw := tar.NewWriter(gw)

		releaseManifest := bytes.NewBuffer([]byte(`---
name: release
version: 1.2.3
`))

		header := &tar.Header{
			Name:    "./release.MF",
			Size:    int64(releaseManifest.Len()),
			Mode:    int64(0644),
			ModTime: time.Now(),
		}

		err := tw.WriteHeader(header)
		Expect(err).NotTo(HaveOccurred())

		_, err = io.Copy(tw, releaseManifest)
		Expect(err).NotTo(HaveOccurred())

		err = tw.Close()
		Expect(err).NotTo(HaveOccurred())

		err = gw.Close()
		Expect(err).NotTo(HaveOccurred())

		filesystem.OpenCall.Returns.File = tarball
	})

	Describe("Read", func() {
		It("extracts the release manifest information from the tarball", func() {
			releaseManifest, err := reader.Read("/path/to/release/tarball")
			Expect(err).NotTo(HaveOccurred())
			Expect(releaseManifest).To(Equal(builder.ReleaseManifest{
				Name:    "release",
				Version: "1.2.3",
			}))

			Expect(filesystem.OpenCall.Receives.Path).To(Equal("/path/to/release/tarball"))
		})

		Context("failure cases", func() {
			Context("when the tarball cannot be opened", func() {
				It("returns an error", func() {
					filesystem.OpenCall.Returns.Error = errors.New("failed to open tarball")

					_, err := reader.Read("/path/to/release/tarball")
					Expect(err).To(MatchError("failed to open tarball"))
				})
			})

			Context("when the input is not a valid gzip", func() {
				It("returns an error", func() {
					filesystem.OpenCall.Returns.File = NewBuffer(bytes.NewBuffer([]byte("I am a banana!")))

					_, err := reader.Read("/path/to/release/tarball")
					Expect(err).To(MatchError("gzip: invalid header"))
				})
			})

			Context("when the header file is corrupt", func() {
				It("returns an error", func() {
					tarball := NewBuffer(bytes.NewBuffer([]byte{}))
					gw := gzip.NewWriter(tarball)
					tw := tar.NewWriter(gw)

					err := tw.Close()
					Expect(err).NotTo(HaveOccurred())

					err = gw.Close()
					Expect(err).NotTo(HaveOccurred())
					filesystem.OpenCall.Returns.File = tarball

					_, err = reader.Read("/path/to/release/tarball")
					Expect(err).To(MatchError("could not find release.MF in \"/path/to/release/tarball\""))
				})
			})

			Context("when there is no release.MF", func() {
				It("returns an error", func() {
					tarball := NewBuffer(bytes.NewBuffer([]byte{}))
					gw := gzip.NewWriter(tarball)
					tw := tar.NewWriter(gw)

					releaseManifest := bytes.NewBuffer([]byte(`---
name: release
version: 1.2.3
`))

					header := &tar.Header{
						Name:    "./someotherfile.MF",
						Size:    int64(releaseManifest.Len()),
						Mode:    int64(0644),
						ModTime: time.Now(),
					}

					err := tw.WriteHeader(header)
					Expect(err).NotTo(HaveOccurred())

					_, err = io.Copy(tw, releaseManifest)
					Expect(err).NotTo(HaveOccurred())

					err = tw.Close()
					Expect(err).NotTo(HaveOccurred())

					err = gw.Close()
					Expect(err).NotTo(HaveOccurred())

					filesystem.OpenCall.Returns.File = tarball
					_, err = reader.Read("/path/to/release/tarball")
					Expect(err).To(MatchError("could not find release.MF in \"/path/to/release/tarball\""))
				})
			})

			Context("when the tarball is corrupt", func() {
				It("returns an error", func() {
					tarball := NewBuffer(bytes.NewBuffer([]byte{}))
					gw := gzip.NewWriter(tarball)
					tw := bufio.NewWriter(gw)

					_, err := tw.WriteString("I am a banana!")
					Expect(err).NotTo(HaveOccurred())

					err = tw.Flush()
					Expect(err).NotTo(HaveOccurred())

					err = gw.Close()
					Expect(err).NotTo(HaveOccurred())

					filesystem.OpenCall.Returns.File = tarball
					_, err = reader.Read("/path/to/release/tarball")
					Expect(err).To(MatchError("error while reading \"/path/to/release/tarball\": unexpected EOF"))
				})
			})

			Context("when the release manifest is not YAML", func() {
				It("returns an error", func() {
					tarball := NewBuffer(bytes.NewBuffer([]byte{}))
					gw := gzip.NewWriter(tarball)
					tw := tar.NewWriter(gw)

					releaseManifest := bytes.NewBuffer([]byte(`%%%%%`))

					header := &tar.Header{
						Name:    "./release.MF",
						Size:    int64(releaseManifest.Len()),
						Mode:    int64(0644),
						ModTime: time.Now(),
					}

					err := tw.WriteHeader(header)
					Expect(err).NotTo(HaveOccurred())

					_, err = io.Copy(tw, releaseManifest)
					Expect(err).NotTo(HaveOccurred())

					err = tw.Close()
					Expect(err).NotTo(HaveOccurred())

					err = gw.Close()
					Expect(err).NotTo(HaveOccurred())

					filesystem.OpenCall.Returns.File = tarball

					_, err = reader.Read("/path/to/release/tarball")
					Expect(err).To(MatchError("yaml: could not find expected directive name"))
				})
			})
		})
	})
})
