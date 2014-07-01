var taskApp = angular.module("taskApp", []);



taskApp.controller("taskCtrl", function taskCtrl($scope, $http) {
    $scope.User = {"FirstName": "Jokie", "LastName":"McFakernammen"};
    $scope.TaskGroups = [
        {"name": "Uncategorized",
         "tasks": [
             {"taskId": 4,
              "taskOwner": 2,
              "taskText":"Get some laundry done!"
             },
             {"taskId": 8,
              "taskOwner": 2,
              "taskText":"Get food..."
             }
         ]},
        {"name": "CRAP",
         "tasks": [
             {"taskId": 88,
              "taskOwner": 2,
              "taskText":"CRY ABOUT THE WORLD"
             },
             {"taskId": 88,
              "taskOwner": 2,
              "taskText":"askd a;skd a;klwmw dasd, asdklnm lkfnoasihf laksdjl kjasldkam lk"
             },
             {"taskId": 88,
              "taskOwner": 2,
              "taskText":"WE CANT SEE ALL THIS, ITS TOO LONG!!!!! CRY ABOUT THE WORLD SAD IJNQEIJ#)( J)(EWFJ )( FJ)(QEFJW)(OE FJPOEWIFJ WPOEIFE "
             },
             {"taskId": 76,
              "taskOwner": 2,
              "taskText":"use fewer \"...\"  ..."
             }
         ]}
    ];

    $scope.FocusedGroup = $scope.TaskGroups[0];

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
