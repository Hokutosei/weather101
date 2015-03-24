'use strict';

var app = angular.module('web102', ['ngRoute', 'ngResource', 'react']);

app.config(function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'js/app/index/template/' + 'index.html',
			controller: 'IndexController'
		})
		.otherwise({
			redirectTo: '/'
		})

	$locationProvider.html5Mode({
		enabled: false,
		requireBase: false
	})

})