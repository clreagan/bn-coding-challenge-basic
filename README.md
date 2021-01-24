# Go Programming Challenge from Buffalo News Challenge Repo
##
# Go API server with goroutines
#
#
# Level -   Advanced
# Time  -   1 - 4 Hours
#
#
# Objective
##

This task is designed to test your knowledge of Go and your ability to problem solve and troubleshoot some areas that may be new to you.

##
# Task
##

Create a GO API server that has a route that uses goroutines to sort numbers in ascending order. The route should run a given number of goroutines, each routine should pause for 1 second and then add to the channel before finishing. Once all the routines have finished the API should return a success message with the list of numbers in the correct order.

The API should:
    1) Reject all other endpoints gracefully
    2) Only accept POST requests
    3) The body of the request should be JSON and look like ```{"numbers": [#,#,#]}``` where # is an arbitrary number.
    4) It should produce an error message if the count of the array of numbers is less than 2 or greater than 100
    5) Use channels
    6) Use goroutines


##
# Notes
##

Have some fun with this and make it your own.
Include lots of comments so that we can get a feel for your thought process while working through this task.
