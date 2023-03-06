$(document).ready(function () {
   // 侧边导航
   (function ($) {

      "use strict";

      var fullHeight = function () {

         $('.js-fullheight').css('height', $(window).height());
         $(window).resize(function () {
            $('.js-fullheight').css('height', $(window).height());
         });

      };
      fullHeight();

      // 隐藏与显示
      $('#sidebarCollapse').on('click', function () {
         $('#sidebar').toggleClass('active');
      });

   })(jQuery);

   initTable();
}); // page loading

var $table = $('#table')
var $add = $('#add')
var $import = $('#import')
var $export = $('#export')
var $remove = $('#remove')

var selections = []

function getIdSelections() {
   return $.map($table.bootstrapTable('getSelections'), function (row) {
      return row.id
   })
}

function responseHandler(res) {
   $.each(res.rows, function (i, row) {
      row.state = $.inArray(row.id, selections) !== -1
   })
   return res
}

function detailFormatter(index, row) {
   var html = []
   $.each(row, function (key, value) {
      html.push('<p><b>' + key + ':</b> ' + value + '</p>')
   })
   return html.join('')
}

function operateFormatter(value, row, index) {
   return [
      '<a class="like" href="javascript:void(0)" title="Like">',
      '<i class="fa fa-heart"></i>',
      '</a>  ',
      '<a class="remove" href="javascript:void(0)" title="Remove">',
      '<i class="fa fa-trash"></i>',
      '</a>'
   ].join('')
}

window.operateEvents = {
   'click .like': function (e, value, row, index) {
      alert('You click like action, row: ' + JSON.stringify(row))
   },
   'click .remove': function (e, value, row, index) {
      $table.bootstrapTable('remove', {
         field: 'id',
         values: [row.id]
      })
      $remove.prop('disabled', !$table.bootstrapTable('getSelections').length)
      $export.prop('disabled', !$table.bootstrapTable('getSelections').length)
   }
}

function totalTextFormatter(data) {
   return 'Total'
}

function totalNameFormatter(data) {
   return data.length
}

function totalPriceFormatter(data) {
   var field = this.field
   return '$' + data.map(function (row) {
      return +row[field].substring(1)
   }).reduce(function (sum, i) {
      return sum + i
   }, 0)
}

function initTable() {
   $table.bootstrapTable('destroy').bootstrapTable({
      height: 550,
      columns: [
         [{
            field: 'state',
            checkbox: true,
            rowspan: 2,
            align: 'center',
            valign: 'middle'
         }, {
            title: 'Item ID',
            field: 'id',
            rowspan: 2,
            align: 'center',
            valign: 'middle',
            sortable: true,
            footerFormatter: totalTextFormatter
         }, {
            title: 'Item Detail',
            colspan: 3,
            align: 'center'
         }],
         [{
            field: 'name',
            title: 'Item Name',
            sortable: true,
            footerFormatter: totalNameFormatter,
            align: 'center'
         }, {
            field: 'price',
            title: 'Item Price',
            sortable: true,
            align: 'center',
            footerFormatter: totalPriceFormatter
         }, {
            field: 'operate',
            title: 'Item Operate',
            align: 'center',
            clickToSelect: false,
            events: window.operateEvents,
            formatter: operateFormatter
         }]
      ]
   })

   $add.prop('disabled', false);
   $import.prop('disabled', false);

   $table.on('check.bs.table uncheck.bs.table ' +
      'check-all.bs.table uncheck-all.bs.table',
      function () {
         $remove.prop('disabled', !$table.bootstrapTable('getSelections').length)
         $export.prop('disabled', !$table.bootstrapTable('getSelections').length)

         // save your data, here just save the current page
         selections = getIdSelections()
         // push or splice the selections if you want to save all data selections
      })
   $table.on('all.bs.table', function (e, name, args) {
      console.log(name, args)
   })

   $remove.click(function () {
      var ids = getIdSelections()
      $table.bootstrapTable('remove', {
         field: 'id',
         values: ids
      })
      $remove.prop('disabled', true)
   })

   $export.click(function () {
      var ids = getIdSelections()
      $export.prop('disabled', true)
   }) // function initTable

   function initTable2() {
      $('#table').bootstrapTable({
         // pagination: true, // 显示记录统计
         // sortable: true,  //是否启用排序
         // sortOrder: "des",  //排序方式 asc升序
         // pageSize: 50, // 每页的记录行数
         // // clickToSelect: true, //是否启用点击选中行
         // uniqueId: "ids",  //每一行的唯一标识，一般为主键列
         // // cardView: true,  //是否显示详细视图
         columns: [{
            field: 'ids',
            title: '序号'
         }, {
            field: 'name',
            title: '姓名'
         }, {
            field: 'gender',
            title: '性别'
         }, {
            field: 'age',
            title: '年龄'
         }, {
            field: 'work_time',
            title: '工作时间'
         }, {
            field: 'salary',
            title: '工资'
         }, {
            field: 'post',
            title: '岗位'
         }, {
            field: 'social_security',
            title: '社保'
         }, {
            field: 'phone',
            title: '电话'
         }, {
            field: 'former_employer',
            title: '原单位'
         }, {
            field: 'height',
            title: '身高'
         }, {
            field: 'weight',
            title: '体重'
         }, {
            field: 'diploma',
            title: '文化'
         }, {
            field: 'political_status',
            title: '政治面貌'
         }, {
            field: 'id',
            title: '身份证'
         }, {
            field: 'security_card',
            title: '保安证'
         }, {
            field: 'current_address',
            title: '现住址'
         }, {
            field: 'comments',
            title: '需要了解的情况'
         }],
         data: [{
            ids: 1,
            name: '张三',
            gender: '男',
            age: '28',
            work_time: '5年',
            salary: '4500',
            post: '保安',
            social_security: '有',
            phone: '13888888888',
            former_employer: '新盾',
            height: '179cm',
            weight: '67kg',
            diploma: '大专',
            political_status: '党员',
            id: '412345678908172844',
            security_card: '123456789',
            current_address: '广西省桂林市七星区五象街道18号',
            comments: '有此情况需要了解'
         }, {
            ids: 2,
            name: '李四',
            gender: '男',
            age: '30',
            work_time: '7年',
            salary: '4800',
            post: '保安',
            social_security: '有',
            phone: '13999999999',
            former_employer: '蓝保',
            height: '181cm',
            weight: '68kg',
            diploma: '大专',
            political_status: '党员',
            id: '412345678908172844',
            security_card: '123456789',
            current_address: '广西省桂林市七星区五象街道18号',
            comments: '有此情况需要了解'
         }, {
            ids: 3,
            name: '王五',
            gender: '男',
            age: '29',
            work_time: '6年',
            salary: '3980',
            post: '保安',
            social_security: '有',
            phone: '13999999999',
            former_employer: '蓝保',
            height: '176cm',
            weight: '66kg',
            diploma: '大专',
            political_status: '党员',
            id: '412345678788172844',
            security_card: '123456711',
            current_address: '广西省桂林市七星区五象街道18号',
            comments: '有此情况需要了解'
         }]
      })
   }
}
