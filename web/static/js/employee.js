$(document).ready(function () {
   $('#table').bootstrapTable({
      columns: [{
         field: 'id',
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
      }],
      data: [{
         id: 1,
         name: '张三',
         gender: '男',
         age: '19'
      }, {
         id: 2,
         name: '李四',
         gender: '女',
         age: '22'
      }]
   })
});
