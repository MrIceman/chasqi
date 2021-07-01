package agent

import "chasqi/rules"

func nodesToRoute(n []rules.Node) *Route {
	rootRoute := Route{previous: nil}
	for idx, item := range n {
		if idx == 0 {
			rootRoute.level = 0
			rootRoute.exposesResponseValues = item.ExposesResponse
			rootRoute.method = item.Method
			rootRoute.name = item.Name
			rootRoute.body = item.RequestBody
			rootRoute.headers = item.RequestHeaders
			rootRoute.path = item.Path
		} else {
			precedentRoute := &rootRoute
			for precedentRoute.level != idx-1 {
				precedentRoute = precedentRoute.next
			}
			route := Route{}
			route.level = idx
			route.exposesResponseValues = item.ExposesResponse
			route.method = item.Method
			route.name = item.Name
			route.headers = item.RequestHeaders
			route.path = item.Path
			route.body = item.RequestBody
			route.previous = precedentRoute
			precedentRoute.next = &route
		}
	}

	return &rootRoute
}
