(function() {
    'use strict';

    app.service('indexService', function($http) {
        return {
            getIndexData: function(callback) {
                $http.get('/get_index').success(callback)
            },
            getCityList: function(callback) {
                $http.get('/get_admin').success(callback)
            }
        }
    })
}());
