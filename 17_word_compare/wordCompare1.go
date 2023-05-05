package wordCompare

import (
	"fmt"
	"strings"
	"sync"
)

func WordCompare1() error {
	textList := []string{
		"IS 616/IEC 60065",
		"トレー·フック",
		"Bluetooth",
		"SGS",
		"33",
		"紙",
		"外箱",
		"内袋·シュリンク",
		"R-41178276",
		"www.bis.gov.in",
		"PAP 21",
		"ETU",
		"WARNING: Risk of reproductive",
		"DAC",
		"IC:23451-A3033C",
		"DE",
		"FCC ID: 2AOKB-A3033C",
		"harm from exposure to Bisphenol",
		"TRI",
		"Raccolta Carta",
		"BANQUETTE",
		"A-www.P65Warnings.ca.gov.",
		"Model: A3033",
		"CAnker Innovations Limited. All rights reserved, registered in the UnitedStates and other countries",
		"Anker Innovations Limited/Room 1318-19, Hollywood Plaza, 610 Nathan Road, Mongkok, Kowloon,",
		"Hong Kong",
		"Anker Technology (UK) Limited/GNR8, 49 Clarendon Road, Watford, Hertfordshire, WD17 1HP,",
		"194644098216",
		"Soundcore Q10I Black",
		"United Kingdom",
		"Anker Innovations Deutschland GmbH/Georg-Muche-Strasse 3, 80807 Munich, Germany",
		"A3033Y11",
		"New",
		"The Bluetooth word mark and logos are registered trademarks owned bytthe Bluetooth SIG, Inc.",
		"AEM-1348",
		"Made in China",
		"and any use of such marks by Anker Innovations Limited is under licenseOther trademarks and",
		"trade names are those of their respective owners",
		"This device complies with Part 15 of the FCC Rules. Operation is subject towing two conditions: (1) This device may not cause harmful",
		"interference, and (2) this device must accept any interference receivedincluding interference that may cause undesired operation",
		"& +1 (800) 988 7973 (US/CA)",
		"& service@soundcore.com (Global)",
		"Discover more at www.soundcore.com",
	}

	targetFlag := make(map[string]bool, 0)
	targetFlag["FCCID"] = false
	targetFlag["IC"] = false
	targetFlag["MODEL"] = false
	targetFlag["A3033Y11"] = false
	targetFlag["AEM-1348"] = false
	targetFlag["194644098216"] = false
	targetFlag["SOUNDCOREQ10IBLACK"] = false

	var targetMap1 sync.Map
	targetMap1.Store("FCCID", "2AOKB-A3033C")
	targetMap1.Store("IC", "23451-A3033C")
	targetMap1.Store("MODEL", "A3033")
	var targetMap2 sync.Map
	targetMap2.Store("194644098216", false)
	targetMap2.Store("SOUNDCOREQ10IBLACK", false)
	targetMap2.Store("A3033Y11", false)
	targetMap2.Store("AEM-1348", false)

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
		if v {
			correct++
		}
	}
	fmt.Println("---------1: ", correct, " / ", total)

	return nil
}
