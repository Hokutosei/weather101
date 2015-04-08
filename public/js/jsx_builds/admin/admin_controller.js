(function() {
    'use strict';
    var log = function(str) { console.log(str) }


    app.controller('AdminIndexController', ['$scope', 'AdminService', function($scope, AdminService) {
        AdminService.indexData(function(data, status) {
            $scope.cityList = data.Data;
        })



    }])


}());
