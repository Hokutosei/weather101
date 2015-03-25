(function() {
	'use strict';
	var log = function(str) { console.log(str); };

	app.controller('IndexController', ['$scope', 'indexService', function($scope, indexService) {
		log("IndexController")
        indexService.getIndexData(function(data, status) {
            log(status)
            log(data)
        })
	}])
}())