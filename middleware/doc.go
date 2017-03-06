/*
Package middleware offers pluggable capabilities of varying purpose such as
session authentication checks. They can be placed in the call stack of any
HTTP handler routine so that they are called prior to advancing to the main
work of each handler function, keeping them small.
*/
package middleware
