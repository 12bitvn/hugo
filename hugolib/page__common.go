// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"github.com/gohugoio/hugo/helpers"
	"sync"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/compare"
	"github.com/gohugoio/hugo/lazy"
	"github.com/gohugoio/hugo/navigation"
	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/resources/resource"
)

type pageCommon struct {
	s *Site
	m *pageMeta

	// Laziliy initialized dependencies.
	init *lazy.Init

	// All of these represents the common parts of a page.Page
	maps.Scratcher
	navigation.PageMenusProvider
	page.AuthorProvider
	page.PageRenderProvider
	page.AlternativeOutputFormatsProvider
	page.ChildCareProvider
	page.FileProvider
	page.GetPageProvider
	page.GitInfoProvider
	page.InSectionPositioner
	page.OutputFormatsProvider
	page.PageMetaProvider
	page.Positioner
	page.RawContentProvider
	page.RelatedKeywordsProvider
	page.RefProvider
	page.ShortcodeInfoProvider
	page.SitesProvider
	page.DeprecatedWarningPageMethods
	page.TranslationsProvider
	page.TreeProvider
	resource.LanguageProvider
	resource.ResourceDataProvider
	resource.ResourceMetaProvider
	resource.ResourceParamsProvider
	resource.ResourceTypesProvider
	resource.TranslationKeyProvider
	compare.Eqer

	// Describes how paths and URLs for this page and its descendants
	// should look like.
	targetPathDescriptor page.TargetPathDescriptor

	layoutDescriptor     output.LayoutDescriptor
	layoutDescriptorInit sync.Once

	// The parsed page content.
	pageContent

	// Set if feature enabled and this is in a Git repo.
	gitInfo *helpers.GitPageInfo

	// Positional navigation
	posNextPrev        *nextPrev
	posNextPrevSection *nextPrev

	// Menus
	pageMenus *pageMenus

	// Internal use
	page.InternalDependencies

	// The children. Regular pages will have none.
	pages     page.Pages
	pagesInit sync.Once

	// Any bundled resources
	resources     resource.Resources
	resourcesInit sync.Once

	translations    page.Pages
	allTranslations page.Pages

	// Calculated an cached translation mapping key
	translationKey     string
	translationKeyInit sync.Once

	// Will only be set for sections and regular pages.
	parent *pageState

	// Will only be set for section pages and the home page.
	subSections page.Pages

	// Set in fast render mode to force render a given page.
	forceRender bool
}
