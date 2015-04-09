(function() {
	'use strict';
	var log = function(str) { console.log(str); };

	app.controller('IndexController', ['$scope', 'indexService', '$timeout', function($scope, indexService, $timeout) {
		$scope.didSearched = false
		$scope.city_data = { Data: [] }

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
		//init()
		$scope.messages = []
		var conn = new WebSocket("ws://localhost:8000/get_index");
		// called when the server closes the connection
		conn.onclose = function(e) {
			$scope.$apply(function(){
				$scope.messages.push("DISCONNECTED");
			});
		};
		// called when the connection to the server is made
		conn.onopen = function(e) {
			$scope.$apply(function(){
				$scope.messages.push("CONNECTED");
			})
		};
		// called when a message is received from the server
		conn.onmessage = function(e){
			//log(e)
			$timeout(function(){
				var data = JSON.parse(e.data)
				log(data)
				$scope.city_data.Data.push(data)
			}, 0);
		};

	}])
}())
