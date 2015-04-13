(function() {
	'use strict';
	var log = function(str) { console.log(str); };

	app.controller('IndexController', ['$scope', 'indexService', '$timeout', function($scope, indexService, $timeout) {
		$scope.didSearched = false
		$scope.city_list = []
		$scope.city_data = { Data: [] }

		$scope.dataStatus = 'Loading...'

		var city_data_init = []

		var init = function(){
			if($scope.didSearched == true) {
				return false
			}

			indexService.getCityList(function(data) {
				$scope.city_list = data.Data.city_list
				log("debug citylist ---")
			})

		}
		init()
		$scope.messages = []
		var host = window.location.hostname
		var conn = new WebSocket("ws://"+ host +"/get_index");
		
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
			var data = JSON.parse(e.data)
			city_data_init.push(data)
			log("receiving...")
			$scope.dataStatus = 'data receiving...'

			if($scope.city_list.length == city_data_init.length) {
				log("will update UI!")
				$scope.$apply(function() {
					$scope.city_data.Data = city_data_init
					log($scope.city_data.Data)
					$scope.dataStatus = 'data finish receiving..'
					city_data_init = []
				})
			}
		};

	}])
}())
