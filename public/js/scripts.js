Vue.component('site-item', {
    data: function () {
        return {
            sitename: "facebook.com",
            status: "200"
        }
    },
    template: '  <div class="col-md-4">' +
        '                    <div class="card mb-4 shadow-sm">' +
        '                        <div class="card-body">' +
        '                            <p class="card-text">{{sitename}}<span class="ml-2 p-1 available-mark green"></span></p>' +
        '                            <div class="d-flex justify-content-between align-items-center">' +
        '                                <small class="text-muted">{{status}}</small>' +
        '                            </div>' +
        '                        </div>' +
        '                    </div>' +
        '                </div>'
})
var app = new Vue({
    el: "#app"
})