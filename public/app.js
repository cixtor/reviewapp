
/* global Vue */

var Application = function () {}

Application.prototype.el = '#reviewapp';

Application.prototype.delimiters = ['${', '}'];

Application.prototype.data = {
    reviews: []
};

Application.prototype.created = function() {
    /**
     * Disregard VueRouter
     *
     * Apparently there is a plugin for Vue called VueRouter that can be
     * attached to the "router" property of the Vue instance to control the
     * application routes and to query the GET parameters. However, after
     * +30 minutes reading the documentation and dabbling with different
     * versions of the code, I was not able to make it work, and since Go
     * can already handle routes very well I suppose there is no point on
     * trying to make this part of the code fancy.
     *
     * @example console.log(this.$route.query.test);
     */
    var app = this;
    var uid = this.queryURL('uid');

    fetch('/reviews/list?uid=' + uid, {method: 'GET'}).then(function (res) {
        if (!res.ok) {
            console.log(res.status + '\x20' + res.statusText + '\x20-\x20' + res.url);
            return;
        }

        res.json().then(function (body) {
            for (var key in body.data) {
                if (body.data.hasOwnProperty(key)) {
                    console.log(body.data[key])
                    body.data[key].cssScore = 'review-stars-score';
                    body.data[key].cssScoreNum = 'review-stars-score-' + body.data[key].score;
                    app.reviews.push(body.data[key]);
                }
            }
        });
    }).catch(function (err) {
        console.log('reviews.uid', err);
    });
};

Application.prototype.methods = {
    queryParams: function () {
        var params = {};
        var parts = [];
        var parameters = location.search.replace(/\?/, '').split('&');

        for (var param in parameters) {
            if (parameters.hasOwnProperty(param)) {
                parts = parameters[param].split('=');
                params[parts[0]] = parts[1];
            }
        }

        return params;
    },

    queryURL: function (query) {
        /**
         * Hardcodes product unique identifier.
         *
         * Since the purpose of this coding challenge is to demonstrate my
         * skills in Go more than any other technology (although it was
         * suggested to use a JavaScript framework) I am hardcoding this UID as
         * the identifier for the product that is going to be reviewed. In a
         * product ready application, this page should not be loaded at all as
         * the Index, we are assuming that the ID will come through a GET
         * parameter but in this example we don't have access to the interface
         * that will be sending the request, so we can only speculate what data
         * will actually be coming, lets hardcode this piece of data for
         * simplicity.
         *
         * @var string
         */
        if (query === 'uid') {
            return 'F79MEIM7';
        }

        var params = this.queryParams();

        if (params[query] === undefined) {
            return null;
        }

        return params[query];
    },

    jsonToQuery: function (json) {
        return Object.keys(json).map(function(key) {
            return encodeURIComponent(key) + '=' +
            encodeURIComponent(json[key]);
        }).join('&');
    },

    showReviewForm: function () {
        document.getElementById('reviewlist').classList.add('d-none');
        document.getElementById('reviewform').classList.remove('d-none');
        document.getElementById('reviewform-alert').classList.add('d-none');
        document.getElementById('reviewform-form').classList.remove('d-none');
    },

    hideReviewForm: function () {
        document.getElementById('reviewform').classList.add('d-none');
        document.getElementById('reviewlist').classList.remove('d-none');
        document.getElementById('reviewform-alert').classList.add('d-none');
        document.getElementById('reviewform-form').classList.remove('d-none');
    },

    submitReview: function (event) {
        var data = {};

        data.uid = this.queryURL('uid');

        for (var key in event.target.elements) {
            if (event.target.elements.hasOwnProperty(key)) {
                if (event.target.elements[key].name !== '') {
                    data[event.target.elements[key].name] = event.target.elements[key].value;
                }
            }
        }

        fetch('/reviews/save', {
            method: 'POST',
            body: this.jsonToQuery(data),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/x-www-form-urlencoded',
            }
        }).then(function (res) {
            if (res.ok) {
                document.getElementById('reviewform-form').classList.add('d-none');
                document.getElementById('reviewform-alert').classList.remove('d-none');

                res.json().then(function (body) {
                    document.getElementById('reviewform-alert')
                    .innerHTML = '<p>' + body.msg + '</p>';
                });
            }
        });
    }
};

new Vue(new Application());
