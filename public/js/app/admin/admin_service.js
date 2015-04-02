(function() {
    'use strict';

    app.service('AdminService', function($http) {
        var api = {
            indexData: function(callback) {
                $http.get('/get_admin').success(callback)
            }
        }
        return api;
    })
}())
