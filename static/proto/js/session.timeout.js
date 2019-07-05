/*
Name: 			Pages / Session Timeout - Examples
Written by: 	Okler Themes - (http://www.okler.net)
Theme Version: 	2.0.0
*/

(function ($) {

    'use strict';

    var SessionTimeout = {

        options: {
            keepAliveUrl: '/proto/session',
            alertOn: 600000, // ms
            timeoutOn: 620000 // ms
        },

        alertTimer: null,
        timeoutTimer: null,

        initialize: function () {
            this
                .start()
                .setup();
        },

        start: function () {
            var _self = this;

            this.alertTimer = setTimeout(function () {
                _self.onAlert();
            }, this.options.alertOn);

            this.timeoutTimer = setTimeout(function () {
                _self.onTimeout();
            }, this.options.timeoutOn);

            return this;
        },

        // bind success callback for all ajax requests
        setup: function () {
            var _self = this;

            // if server returns successfuly,
            // then the session is renewed.
            // thats why we reset here the counter
            $(document).ajaxSuccess(function () {
                _self.reset();
            });

            return this;
        },

        reset: function () {
            clearTimeout(this.alertTimer);
            clearTimeout(this.timeoutTimer);
            this.start();

            return this;
        },

        keepAlive: function () {
            // we don't have session on demo,
            // so the code above prevent a request to be made
            // in your project, please remove the next 3 lines of code
            // if (!this.options.keepAliveUrl) {
            //     this.reset();
            //     return;
            // }

            var _self = this;

            $.ajax({
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                url: this.options.keepAliveUrl,
                type: 'POST',
                data: {
                },
                dataType: 'json',
                success: function (data) {
                    _self.reset();
                },
            })
        },

        // ------------------------------------------------------------------------
        // CUSTOMIZE HERE
        // ------------------------------------------------------------------------
        onAlert: function () {
            // if you want to show some warning
            // TODO: remove this confirm (it breaks the logic and it's ugly)
            var renew = confirm('您还在电脑前吗? 账户将在20秒钟后自动退出登录');

            if (renew) {
                this.keepAlive();
            }

            // if you want session to not expire
            // this.keepAlive();
        },

        onTimeout: function () {
            self.location.href = '/logout';
        }

    };

    $(function () {
        SessionTimeout.initialize();
    });

}).apply(this, [jQuery]);
