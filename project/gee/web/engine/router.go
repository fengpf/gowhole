package engine

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

//func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
//	log.Printf("addRoute route %4s - %s", method, pattern)
//	key := method + "-" + pattern
//
//	r.handlers[key] = handler
//}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("addRoute route %4s - %s", method, pattern)

	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler

	//spew.Dump(r.roots)
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)

	params := make(map[string]string)
	root, ok := r.roots[method]

	//spew.Dump(method, path, searchParts, root)

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}

		return n, params
	}

	return nil, nil
}

//func (r *router) handle(c *Context) {
//	key := c.Method + "-" + c.Path
//
//	if handler, ok := r.handlers[key]; ok {
//		handler(c)
//	} else {
//		c.Plain(http.StatusNotFound, "404 Not Found:%s\n", c.Path)
//	}
//}

//func (r *router) handle(c *Context) {
//	n,params:=r.getRoute(c.Method, c.Path)
//
//	if n!=nil{
//		c.Params  = params
//		key := c.Method + "-" + n.pattern
//		r.handlers[key](c)
//	} else {
//		c.Plain(http.StatusNotFound, "404 Not Found:%s\n", c.Path)
//	}
//}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.Plain(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next()
}
