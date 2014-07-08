var ListApp = angular.module("ListApp", []);

ListApp.controller('ListCtrl', ['$scope', '$http', function($scope, $http) {
    console.log("In ListCtrl");

    $http({
        method:"GET",
        url:"/"
    }).success(function(data, status, headers, config) {
        //$scope.Lists = data;
        $scope.Lists = [{title:"Chores",items:["Do the laundry","Clean the catbox"]},
                        {title:"Life goals",items:["Kill all the humans","Eat too much pizza","become a cat"]},
                        {title:"Super cool things to share",items:["http://golang.org/pkg/database/sql/#DB.Prepare","http://golang.org/pkg/crypto/","http://golang.org/pkg/fmt/","http:///"]},
                        {title:"Groceries",items:["Bananas","soap","chips","milk","Pasta","More kitty food"]}];
    }).error(function(data, status, headers, config) {
        console.log("Error getting lists");
    });


}]);
