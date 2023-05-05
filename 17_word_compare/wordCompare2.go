package wordCompare

import (
	"fmt"
	"strings"
	"sync"
)

func WordCompare2() error {
	textList := []string{
		"FR",
		"ETUI",
		"BAC",
		"吃饭",
		"BARQUETTE",
		"紙",
		"内蓋",
		"外箱",
		"FCC ID: 2AOKB-A3028",
		"DAnker Innovations Limited. All nghts reserved registered in the United Stuntres",
		"Anker Innovations Limitediroom 1318-19, Hollywood Plaza, 610 Nathan Road, Hong Kong",
		"Bluetooth",
		"Anker Technology (UK) Limited/GNRS, 49 Clarendon Road, Watfordshint, WD17 ited Kingdom",
		"Anker Innovations Deutschland GmbH/Georg-Muche-Strasse 3, 80807 Munich Germ",
		"The Bluetooth* word mark and logos are registered trademaris owned by the ",
		"Inc.and any use of such marks by Anker Innovations Limited is underlicense",
		"SGS",
		"Other trademarks and trade names are those of their respective owners",
		"IC:23451-A3028",
		"This device complies with Part 15 of the FCC Rules Operation is subject totwo conditions",
		"(1) This device may not cause harmful interference, and (2) his device mus",
		"interference received, including interence that may cause undesired operat ",
		"+1(800)988 7973 (US/CA)",
		"Discover more at www.soundcore.com",
		"Model: A3028",
		"service@soundcore.com",
		"Soundcore",
		"What's In The Box",
		"Made in China",
		"600",
		"USB-C Cable",
		"AUX Cable",
		"Made in China",
		"New",
		"Soundcore Q30 Black - EU ONLINE",
		"Travel Case",
		"X001NSXO27",
		"A3028311",
		"AEM",
		"www.bis.gov.in",
		"R-41178276",
		"Raccolta Carta",
		"PAP 20",
		"IS 616/IEC 60065",
	}

	targetFlag := make(map[string]bool, 0)
	targetFlag["FCCID"] = false
	targetFlag["IC"] = false
	targetFlag["MODEL"] = false
	targetFlag["A3028311"] = false
	targetFlag["AEM"] = false
	targetFlag["X001NSXO27"] = false
	targetFlag["SOUNDCOREQ30BLACK-EUONLINE"] = false

	var targetMap1 sync.Map
	targetMap1.Store("FCCID", "2AOKB-A3028")
	targetMap1.Store("IC", "23451-A3028")
	targetMap1.Store("MODEL", "A3028")
	var targetMap2 sync.Map
	targetMap2.Store("X001NSXO27", false)
	targetMap2.Store("SOUNDCOREQ30BLACK-EUONLINE", false)
	targetMap2.Store("A3028311", false)
	targetMap2.Store("AEM", false)

	for _, s := range textList {
		str := strings.ReplaceAll(s, " ", "")
		str = strings.ToUpper(str)

		targetMap1.Range(func(key, value any) bool {
			k, v := key.(string), value.(string)
			containsK, containsV := strings.Contains(str, k), strings.Contains(str, v)
			if containsK && containsV {
				targetFlag[k] = true
				targetMap1.Delete(k)
			}
			return true
		})
		targetMap2.Range(func(key, value any) bool {
			k := key.(string)
			containsK := strings.Contains(str, k)
			if containsK {
				targetFlag[k] = true
				targetMap2.Delete(k)
			}
			return true
		})
	}

	correct, total := 0, len(targetFlag)
	for _, v := range targetFlag {
		//fmt.Println(k, "\t ", v)
		if v {
			correct++
		}
	}
	fmt.Println("---------2: ", correct, " / ", total)


	return nil
}
