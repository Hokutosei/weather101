'use strict';

var app = angular.module('web102', ['ngRoute', 'ngResource', 'react']);

app.config(function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'js/app/index/template/' + 'index.html',
			controller: 'IndexController'
		})
		.when('/admin', {
			templateUrl: 'js/app/admin/template/' + 'index.html',
			controller: 'AdminIndexController'
		})
		.otherwise({
			redirectTo: '/'
		})

	$locationProvider.html5Mode({
		enabled: true,
		requireBase: true
	})
})
