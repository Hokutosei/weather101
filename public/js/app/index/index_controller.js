(function() {
	'use strict';
	var log = function(str) { console.log(str); };

	app.controller('IndexController', ['$scope', 'indexService', function($scope, indexService) {
        indexService.getIndexData(function(data, status) {
            log(status)
            log(data)
            $scope.city_data = data;
        })
	}])
}())