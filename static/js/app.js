var taskApp = angular.module("taskApp", []);



taskApp.controller("taskCtrl", function taskCtrl($scope, $http) {
    console.log("In taskCtrl");
    $scope.User = {};
    // TODO: Discuss data format and presentation
    $scope.TaskGroups = $http({method: "GET",
                                url: "/task/get/all"
                               }).success(function(data, status, headers, config){
                                   return data.Data;
                               }).error(function(data, status, headers, config){
                                   console.log("ERROR: Could not retrieve tasks.");
                                   // Some kind of flashy div error message. Maybe try again as well?
                                   return {};
                               });

    $scope.FocusedGroup = $scope.TaskGroups[0];

    $scope.TestQuery = function() {
        $http({method: "GET",
               url: "/task/get/all"
              }).success(function(data, status, headers, config){
                  console.log(data);
              }).error(function(data, status, headers, config){
                  console.log("ERROR: Could not retrieve tasks.");
              });
    };

    $scope.UpdateFocusedGroup = function(group) {
        $scope.FocusedGroup = group;
    };

    $scope.Addtask = function() {

    };


    $scope.DeleteTask = function() {

    };


    $scope.UpdateTask = function() {

    };


    $scope.ClearCompleted = function() {

    };


});
