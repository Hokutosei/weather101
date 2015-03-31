(function() {
	'use strict';
	var log = function(str) { console.log(str); };

	app.controller('IndexController', ['$scope', 'indexService', function($scope, indexService) {
		$scope.didSearched = false

		var init = function(){
			if($scope.didSearched == true) {
				return false
			}
			indexService.getIndexData(function(data, status) {
				log(status)
				log(data)
				$scope.didSearched = true
				$scope.city_data = data;
			})
		}
		init()
	}])
}())
