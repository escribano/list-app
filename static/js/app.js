var ListApp = angular.module("ListApp", []);

ListApp.controller('ListCtrl', ['$scope', function($scope) {
    console.log("In ListCtrl");

    $scope.Lists = [
        {
            title:"Chores",
            items:[
                "Do the laundry", "Clean the catbox"
            ]
        },
        {
            title:"Life goals",
            items:[
                "Kill all the humans", "Eat too much pizza", "become a cat"
            ]
        },
        {
            title:"Super cool things to share",
            items:[
                "http://golang.org/pkg/database/sql/#DB.Prepare",
                "http://golang.org/pkg/crypto/",
                "http://golang.org/pkg/fmt/",
                "http://golang.org/pkg/fsnotify/",
                "http://golang.org/pkg/inotify/"
            ]
        },
        {
            title:"Groceries",
            items:[
                "Bananas", "soap", "chips", "milk", "Pasta", "More kitty food"
            ]
        }
    ];

}]);
