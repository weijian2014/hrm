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

      // for debug
      $('#sidebar').toggleClass('active');

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
      '<a class="like" href="javascript:void(0)" title="修改">',
      '<i class="fa fa-file"></i>',
      '</a>  ',
      '<a class="remove" href="javascript:void(0)" title="删除">',
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

function loadingTemplate(message) {
   var type = 'fa'
   if (type === 'fa') {
      return '<i class="fa fa-spinner fa-spin fa-fw fa-2x"></i>'
   }
   if (type === 'pl') {
      return '<div class="ph-item"><div class="ph-picture"></div></div>'
   }
}

function genderFormatter(value, row) {
   var icon = value == "男" ? 'fa fa-mars' : 'fa fa-venus'
   return '<i class="' + icon + '"></i> ' + value
}

function initTable() {
   $table.bootstrapTable('destroy').bootstrapTable({
      url: '../../doc/testData.json',
      // url: 'https://examples.wenzhixin.net.cn/examples/bootstrap_table/data',
      method: 'get',                      //请求方式（*）
      toolbar: '#toolbar',                //工具栏按钮用哪个容器 一个jQuery 选择器，指明自定义的 buttons toolbar。例如:#buttons-toolbar, .buttons-toolbar 或 DOM 节点
      toolbarAlign: 'left',                 //指示如何对齐自定义工具栏。可以使用'left'，'right'
      buttonsToolbar: '',                  //一个jQuery选择器，指示按钮工具栏，例如：＃buttons-toolbar，.buttons-toolbar或DOM节点
      buttonsAlign: 'right',               //指示如何对齐工具栏按钮。可以使用'left'，'right'。
      buttonsClass: 'primary',
      striped: false,                      //是否显示行间隔色
      cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
      pagination: true,                   //是否显示分页（*） 设置为true以在表格底部显示分页工具栏默认false
      paginationLoop: true,               //设置为true以启用分页连续循环模式    
      paginationHAlign: 'right',          //分页条水平方向的位置，默认right（最右），可选left   
      paginationDetailHAlign: 'left',     //如果解译的话太长，举个例子，paginationDetail就是“显示第 1 到第 8 条记录，总共 15 条记录 每页显示 8 条记录”，默认left（最左），可选right    
      paginationVAlign: 'bottom',         //分页条垂直方向的位置，默认bottom（底部），可选top、both（顶部和底部均有分页条）    
      paginationPreText: '上一页',        //上一页的按钮符号
      paginationNextText: '下一页',       //下一页的按钮符号    
      paginationSuccessivelySize: 5,      //分页时会有<12345...80>这种格式而5则表示显示...左边的的页数   
      paginationPagesBySide: 5,           //...右边的最大连续页数如改为2则 <1 2 3 4....79 80>   
      paginationUseIntermediate: false,   //计算并显示中间页面以便快速访问 true 会将...替换为计算的中间页数42
      onlyInfoPagination: false,          //设置为true以仅显示表中显示的数据量, 这会导致分页失效
      showExtendedPagination: true,       // 分页显示的增强, 比如搜索时可以显示 
      totalNotFilteredField: true,        // 从 xx 总记录中过滤
      sortable: true,                     // 全局的, 是否启用排序, 列中也有此变量
      sortName: '',                        //全局的, 可以在列中单独使用sortName: 'name', 定义要排序的列, 没定义默认都不排列，同sortOrder结合使用，sortOrder没写的话列默认递增（asc）
      sortOrder: "asc",                   //全局的, 可以在列中单独使用sortOrder: 'desc', 定义列排序顺序，只能是'asc'或'desc'。
      sortResetPage: true,                // 排序时重置分页信息
      sortStable: false,                   //如果你把此属性设为了true）我们将为此行添加'_position'属性 （别看错了，是sortStable，sortable在下面）设为true，则和sort部分一样，区别是：在排序过程中，如果存在相等的元素，则原来的顺序不会改变
      sidePagination: "client",           //分页方式：client客户端分页（默认），server服务端分页（*）
      silentSort: true,                   //设置为false以便对加载的消息数据进行排序。当sidePagination选项设置为“server”时，此选项有效。 
      pageNumber: 1,                       //初始化加载第一页，默认第一页
      pageSize: 10,                       //每页的记录行数（*）
      pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
      search: true,                       //是否显示表格搜索input，此搜索是客户端搜索，不会进服务端，所以，个人感觉意义不大
      strictSearch: false,                 //启用严格搜索
      searchOnEnterKey: false,            // true时搜索方法将一直执行，直到按下Enter键（即按下回车键才进行搜索）   
      trimOnSearch: true,                 //默认true，自动忽略空格        
      searchAlign: 'right',                //指定搜索输入框的方向。可以使用'left'，'right'。    
      // searchTimeOut: 500,                 //设置搜索触发超时
      searchText: '',                     //设置搜索文本框的默认搜索值
      showSearchButton: false,             // 搜索框搜索按钮
      showSearchClearButton: true,        // 搜索框内容清除按钮
      searchHighlight: true,              //搜索高亮
      showColumns: true,                 //是否显示所有的列 设置为true以显示列下拉列表（一个可以设置显示想要的列的下拉f按钮）
      visibleSearch: false,               // 是否只搜索可见数据
      showRefresh: true,                  //是否显示刷新按钮 默认false
      minimumCountColumns: 5,             //最少允许的列数  要从列下拉列表中隐藏的最小列数
      showColumnsToggleAll: true,         //显示或者隐藏所有列
      clickToSelect: true,                //是否启用点击选中行
      idField: 'id',                      //表明哪个是字段是标识字段
      uniqueId: "id",                     //表明每一行的唯一标识字段，一般为主键列
      showToggle: true,                    //是否显示详细视图和列表视图的切换按钮
      cardView: false,                    //是否显示详细视图  设置为true以显示卡片视图表，例如mobile视图（卡片视图）
      detailView: false,                  //设置为true以显示detail 视图表（细节视图）
      detailViewByClick: false,
      locale: 'zh-CN',
      classes: 'table table-bordered table-hover table-sm',  //表格样式。可以使用'table'，'table-bordered'，'table-hover'，'table-striped'，'table-dark'，'table-sm'和'table-borderless'。默认情况下，表格是有界的。
      theadClasses: '',                      // 表thead的类名 如使用.thead-light或.thead-dark使theads显示为浅灰色或深灰色。
      undefinedText: 'N/A',                    // 定义默认的未定义文本   
      sortClass: '',                         //已排序的td元素的类名
      rememberOrder: false,                     //设置为true以记住每列的顺序
      contentType: 'application/json',          //请求远程数据的contentType，例如：application/x-www-form-urlencoded。    
      dataType: 'json',                //您希望服务器返回的数据类型    
      totalField: 'total',          //Key in incoming json containing 'total' data.
      dataField: 'rows',            //名称写自己定义的每列的字段名，也就是key，通过key才能给某行的某列赋value原文：获取每行数据json内的key   
      totalRows: 0,                 //该属性主要由分页服务器传递，易于使用      
      showHeader: true,          //设置为false以隐藏表头    
      showFooter: true,          //设置为true以显示摘要页脚行(固定也交 比如显示总数什么的最合适)    
      showPaginationSwitch: true,         //设置为true以显示分页组件的切换按钮    
      showFullscreen: true,         // 设置为true以显示全屏按钮
      smartDisplay: true,        //设置为true以巧妙地显示分页或卡片视图
      escape: false,                      // 转义字符串以插入HTML，替换 &, <, >, “, `, 和 ‘字符  跳过插入HTML中的字符串，替换掉特殊字符 
      selectItemName: 'btSelectItem',//  设置radio 或者 checkbox的字段名称
      clickToSelect: true,          //设置为true时 在点击列时可以选择checkbox或radio
      singleSelect: false,           //默认false，设为true则允许复选框仅选择一行
      checkboxHeader: true,            //设置为false以隐藏标题行中的check-all复选框 即隐藏全选框
      maintainSelected: true,         // true时点击分页按钮或搜索按钮时，记住checkbox的选择项 设为true则保持被选的那一行的状态
      showButtonText: true,         //设置表格右上角的图标是否显示文字
      showExport: true,             // 显示导出数据-表格右上角
      // exportDataType : "all",    // 导出所有数据, 否则只导出页面显示的
      exportOptions: {              // https://github.com/hhurz/tableExport.jquery.plugin
         ignoreColumn: [0],
         fileName: '导出数据',
         worksheetName: 'sheet1',
         csvEnclosure: '"',
         csvSeparator: ',',
         csvUseBOM: true,
      },
      icons: {                   //定义工具栏，分页和详细信息视图中使用的图标
         paginationSwitchDown: 'fa-caret-square-down',
         paginationSwitchUp: 'fa-caret-square-up',
         refresh: 'fa-sync',
         toggleOff: 'fa-toggle-off',
         toggleOn: 'fa-toggle-on',
         columns: 'fa-th-list',
         detailOpen: 'fa-plus',
         detailClose: 'fa-minus',
         fullscreen: 'fa-arrows-alt',
         clearSearch: 'fa-undo'
      },
      iconsPrefix: 'fa',//定义图标集名称（FontAwesome的'glyphicon'或'fa'）
      iconSize: 'btn-xssm',// 定义icon图表 undefined => btnxs => btn-xssm => btn-smlg => btn-lg
      responseHandler: responseHandler,
      detailFormatter: detailFormatter,
      loadingTemplate: loadingTemplate,
      columns: [{
         field: 'select',
         title: '选择',
         checkbox: true,
         align: 'center',
         halign: 'center',
         valign: 'middle',
      }, {
         field: 'id',
         title: '序号',
         visible: false,
         sortable: true,
         searchable: false,
         align: 'center',
         valign: 'middle',
      }, {
         field: 'name',
         title: '姓名',
         sortable: true,
         align: 'center'
      }, {
         field: 'gender',
         title: '性别',
         sortable: true,
         align: 'center',
         formatter: genderFormatter,
      }, {
         field: 'age',
         title: '年龄',
         sortable: true,
      }, {
         field: 'work_time',
         title: '工作时间',
         sortable: true,
      }, {
         field: 'salary',
         title: '工资',
         sortable: true,
      }, {
         field: 'post',
         title: '岗位',
      }, {
         field: 'social_security',
         title: '社保',
      }, {
         field: 'phone',
         title: '电话',
      }, {
         field: 'former_employer',
         title: '原单位',
      }, {
         field: 'height',
         title: '身高',
         sortable: true,
      }, {
         field: 'weight',
         title: '体重',
         sortable: true,
      }, {
         field: 'diploma',
         title: '文化',
      }, {
         field: 'political_status',
         title: '政治面貌',
      }, {
         field: 'identifier',
         title: '身份证',
      }, {
         field: 'security_card',
         title: '保安证',
      }, {
         field: 'current_address',
         title: '现住址',
         align: 'left',
         halign: 'center',
         valign: 'center',
      }, {
         field: 'comments',
         title: '需要了解的情况',
         visible: false,
         align: 'left',
         halign: 'center',
         valign: 'center',
      }, {
         field: 'operate',
         title: '操作',
         clickToSelect: false,
         events: window.operateEvents,
         formatter: operateFormatter
      }]
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
   })

   $add.click(function () {
      alert('You click add')
   })

   $import.click(function () {
      alert('You click import')
   })


}  // function initTable2()
