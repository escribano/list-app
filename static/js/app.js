var taskApp = angular.module("taskApp", []);

taskApp.service('Tasks', ['$rootScope', '$http', function($rootScope, $http) {
    var service = {
        groups: {},

        addGroup: function(groupName) {},
        removeGroup: function(groupId) {},

        addTask: function(taskText, taskGroup, taskDueDate, subTasks) {},
        removeTask: function(taskId) {}
    };

    $http({
        method:'GET',
        url: '/task/get/all'
    }).success(function(data, status, headers, config) {
        console.log("TASKS: ", data.data);
    }).error(function(data, status, headers, config) {
        console.log("ERROR: Could not retreive tasks.");
    });

    return service;
}]);

taskApp.controller('taskCtrl', ['Tasks', function taskCtrl($scope, $http, Tasks) {

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


}]);
