$(document).ready(function () {
   (function($) {

      "use strict";
   
      var fullHeight = function() {
   
         $('.js-fullheight').css('height', $(window).height());
         $(window).resize(function(){
            $('.js-fullheight').css('height', $(window).height());
         });
   
      };
      fullHeight();
   
      // 侧边导航
      $('#sidebarCollapse').on('click', function () {
         $('#sidebar').toggleClass('active');
     });
   
   })(jQuery);

   $('#table').bootstrapTable({
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
});
