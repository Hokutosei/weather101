(function() {
    'use strict';

    app.service('indexService', function($http) {
        return {
            getIndexData: function(callback) {
                $http.get('/get_index').success(callback)
            }
        }
    })
}());