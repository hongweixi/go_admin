// (function () {
// $('.date-picker').datepicker({
// 	format: "yyyy-mm-dd",
// 	todayBtn: "linked",
// 	language: "zh-CN",
// 	calendarWeeks: true,
// 	autoclose: true,
// 	todayHighlight: true,
// 	toggleActive: true,
// 	defaultViewDate: {year: new Date().getFullYear(), month: new Date().getMonth(), day: new Date().getd}
// });
var datatableInit = function (table_obj) {
    var obj = table_obj;
    var url = obj.data('url');
    var role = obj.data('role');
    // 1.
    obj.DataTable({
        ordering: false,
        processing: true,
        responsive: true,
        // "scrollX": true,
        lengthMenu: [
            [10, 25, 50, -1],
            ['10', '25', '50', 'All']
        ],
        language: {
            url: 'https://cdn.datatables.net/plug-ins/1.10.16/i18n/Chinese.json'
        },
        stateSave: true,
        "stateSaveParams": function (settings, data) {
            //这里可以操作保存的数据，写上自己特定的逻辑
            //data.search.search = "";
        },
        "stateSaveCallback": function (settings, data) {
            sessionStorage.setItem(role + settings.sInstance, JSON.stringify(data));
        },
        //读取状态操作
        "stateLoadParams": function (settings, data) {
            //在读取数据的时候可以改变数据，根据自己逻辑来处理
            //data.search.search = "";
            //或者你可以直接禁用从缓存里读取数据，只要直接返回false即可
            //return false;
        },
        "stateLoadCallback": function (settings) {
            return JSON.parse(sessionStorage.getItem(role + settings.sInstance));
        },
        //状态加载完后执行的回调函数

        "stateLoaded": function (settings, data) {
            for (var i = 0; i < data.columns.length; i++) {
                if (data.columns[i].search.search != "") {
                    $("." + role + "_" + i).val(data.columns[i].search.search);
                }
            }
            //在这里你可以打印出保存的缓存数据
            //alert( 'Saved filter was: '+data.search.search );
        },
        serverSide: true,
        dom: 'B<"clear">lrtip',
        buttons: ['excel', 'print', 'pdf', 'copy'],
        ajax: {
            url: url,
            type: 'POST',
            headers: {
                'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
            }
        },
        "deferRender": true
    });
}

// 2



function InitDatatable(table_role) {
    var table_obj = $('#' + table_role + "_table");

    datatableInit(table_obj)
    $('.' + table_role + '_column_filter').on('change', function () {
        table_obj.DataTable().column($(this).data('column')).search($(this).val(), false, true).draw();
    });
}

// $('.column_filter').on('change', function () {
// 	obj.DataTable().column($(this).data('column')).search($(this).val(), false, true).draw();
// });
// })();