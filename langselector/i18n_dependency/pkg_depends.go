/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package i18n_dependency

import (
	libsoft "dbus/com/linuxdeepin/softwarecenter"
	"pkg.linuxdeepin.com/dde-daemon/langselector/language_info"
	"strings"
)

const (
	formatTypeNone int32 = 0
	//Format: %LCODE%
	formatTypeLC int32 = 1
	//Format: %LCODE% or %LCODE%-%CCODE%
	formatTypeLCCC int32 = 2
	//Format: %LCODE% or %LCODE%%CCODE% or %LCODE%-%VARIANT%
	formatTypeLCVA int32 = 3
)

const (
	I18nDependencyFilename = "/usr/share/dde-daemon/lang/pkg_depends.json"
	LanguageListFile       = "/usr/share/dde-daemon/lang/support_languages.json"
)

type packagesDepend struct {
	depend   string
	packages []string
}

func InstallDependentPackages(locale, pkgConfig, i18nConfig string) error {
	softCenter, err := libsoft.NewSoftwareCenter(
		"com.linuxdeepin.softwarecenter",
		"/com/linuxdeepin/softwarecenter")
	if err != nil {
		return err
	}

	dependsList, err := getPkgDependList(pkgConfig)
	if err != nil {
		return err
	}

	var packages []string
	trList := getDependentPkgListByKey("tr", locale,
		i18nConfig, dependsList)
	list := getPkgListFromPackagesDependList(trList, softCenter)
	waList := getDependentPkgListByKey("wa", locale,
		i18nConfig, dependsList)
	packages = append(packages, list...)
	list = getPkgListFromPackagesDependList(waList, softCenter)
	packages = append(packages, list...)

	err = softCenter.InstallPkg(packages)
	return err
}

func isPkgInstalled(pkg string, softCenter *libsoft.SoftwareCenter) bool {
	if len(pkg) == 0 {
		return true
	}

	ret, _ := softCenter.GetPkgInstalled(pkg)
	if ret == 1 {
		return true
	}

	return false
}

func isPkgExist(pkg string, softCenter *libsoft.SoftwareCenter) bool {
	if len(pkg) == 0 {
		return false
	}

	list, err := softCenter.IsPkgInCache(pkg)
	if err != nil {
		return false
	}

	for _, p := range list {
		if pkg == p {
			return true
		}
	}

	return false
}

func getPkgListFromPackagesDependList(infoList []packagesDepend,
	softCenter *libsoft.SoftwareCenter) []string {
	var packages []string
	for _, info := range infoList {
		if !isPkgInstalled(info.depend, softCenter) {
			continue
		}

		for _, pkg := range info.packages {
			if !isPkgExist(pkg, softCenter) {
				continue
			}
			packages = append(packages, pkg)
		}
	}

	return packages
}

func getDependentPkgListByKey(key, locale, config string,
	dependsList *dependentPkgList) (packages []packagesDepend) {
	var pkgInfoList []dependentPkgInfo
	for _, group := range dependsList.PkgDepends {
		if key == group.Category {
			pkgInfoList = group.PkgInfos
			break
		}
	}

	for _, info := range pkgInfoList {
		if info.LangCode == "" {
			list := parseFormatType(&info, locale, config)
			tmp := packagesDepend{info.DependPkg, list}
			packages = append(packages, tmp)
		} else {
			lcode, _, _ := getLangCodeByLocale(locale, config)
			if lcode == info.LangCode {
				tmp := packagesDepend{info.DependPkg, []string{info.PkgPull}}
				packages = append(packages, tmp)
			}
		}
	}

	return
}

func parseFormatType(info *dependentPkgInfo, locale, config string) (pkgList []string) {
	if info == nil {
		return
	}

	switch info.FormatType {
	case formatTypeNone:
		pkgList = append(pkgList, info.PkgPull)
	case formatTypeLC:
		lcode, _, _ := getLangCodeByLocale(locale, config)
		if len(lcode) > 0 {
			pkg := info.PkgPull + lcode
			pkgList = append(pkgList, pkg)
		}
	case formatTypeLCCC:
		lcode, ccode, _ := getLangCodeByLocale(locale, config)
		ccode = strings.ToLower(ccode)
		if len(lcode) > 0 {
			pkg := info.PkgPull + lcode
			pkgList = append(pkgList, pkg)
		}

		if len(lcode) > 0 && len(ccode) > 0 {
			pkg := info.PkgPull + lcode + "-" + ccode
			pkgList = append(pkgList, pkg)
		}
	case formatTypeLCVA:
		lcode, ccode, variant := getLangCodeByLocale(locale, config)
		ccode = strings.ToLower(ccode)
		if len(lcode) > 0 {
			pkg := info.PkgPull + lcode
			pkgList = append(pkgList, pkg)
		}

		if len(lcode) > 0 && len(ccode) > 0 {
			pkg := info.PkgPull + lcode + ccode
			pkgList = append(pkgList, pkg)
		}

		if len(lcode) > 0 && len(variant) > 0 {
			pkg := info.PkgPull + lcode + "-" + variant
			pkgList = append(pkgList, pkg)
		}
	}

	return
}

func getLangCodeByLocale(locale, config string) (string, string, string) {
	var (
		lcode   string
		ccode   string
		variant string
	)

	info, err := language_info.GetCodeInfoByLocale(locale,
		config)
	if err != nil {
		return lcode, ccode, variant
	}

	lcode = info.LangCode
	ccode = info.CountryCode
	variant = info.Variant

	return lcode, ccode, variant
}