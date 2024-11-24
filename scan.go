package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Definisikan flag untuk menerima argumen --target
	var targetFolder string
	flag.StringVar(&targetFolder, "target", "", "Path ke folder yang ingin diperiksa")
	flag.Parse()

	if targetFolder == "" {
		fmt.Println("Harap tentukan path folder dengan menggunakan flag --target")
		return
	}

	// Kata kunci yang ingin Anda cocokkan
	keywords := []string{
		"787backlink.xyz",
		"binusrun.com",
		"seokeras.com",
		"jasaiklan.lol",
		"umrohkemenag.org",
		"a.mikomallkopo.com",
		"purefine.online",
		"twins-4e372.web.app",
		"forbesstream.com",
		"pinjam100.site",
		"fixtechcs.net",
		"dw88.cwhonors.org",
		"w88.sepangracingteam.com",
		"ant1rungk4d",
		"ampshopify.store",
		"petirjago.site",
		"seoinclude.pages.dev",
		"includes-page.com",
		"wp.rpka.org",
		"admin.rpka.org",
		"rpka.org",
		"hedon77official.xyz",
		"linkgacorthailand.xyz",
		"dp288.live",
		"bobabotui.lol",
		"seobeton.team",
		"aHR0cHM6Ly95b2licmUuY29tL3NoZWxsL2FkbWluLnR4dA",
		"casino",
		"winrate",
		"gacor",
		"togel",
		"maxwin",
		"scatter",
		"nextcloud",
		"owncloud",
	}

	// Lakukan pemindaian pada semua file di folder
	err := filepath.Walk(targetFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Memeriksa apakah file adalah symlink
		if !info.Mode().IsRegular() {
			return nil // Mengabaikan symlink
		}

		if !info.IsDir() {
		// Hanya memeriksa file .php saja
		//if !info.IsDir() && strings.HasSuffix(info.Name(), ".php") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Cocokkan dengan kata kunci
			for _, keyword := range keywords {
				if strings.Contains(string(content), keyword) {
					fmt.Printf("File %s keyword: %s\n", path, keyword)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
	}
}
