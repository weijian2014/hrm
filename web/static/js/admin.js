$(document).ready(function () {
   doPaging(true, 1);

   //点击搜索按钮时
   $('.searchButton').click(function (event) {
      var keyWord = $('input[name="searchInput"]').val();
      if ('' == keyWord) {
         return;
      }

      adminSearch(keyWord);
   });

   //点击新增按钮时
   $('.addBtn').click(function (event) {
      var vehicleName = $('input[name="vehicleNameInput"]').val();
      var partsName = $('input[name="partsNameInput"]').val();
      var partsNo = $('input[name="partsNoInput"]').val();
      var partsCount = Number($('input[name="partsCountInput"]').val());
      var partsPrice = Number($('input[name="partsPriceInput"]').val());
      var partsSellingPrice = Number($('input[name="partsSellingPriceInput"]').val());

      if ('' == vehicleName ||
         '' == partsName ||
         '' == partsNo) {
         printAtMousePointer(this, "输入的必填字段有空值，请检查后重新提交！");
         return;
      }

      if (isNaN(partsCount)) {
         printAtMousePointer(this, "输入的配件个数为非法值，请检查后重新提交！");
         $('input[name="partsCountInput"]').focus();
         return;
      }

      if (isNaN(partsPrice)) {
         printAtMousePointer(this, "输入的配件单价为非法值，请检查后重新提交！");
         $('input[name="partsPriceInput"]').focus();
         return;
      }

      if (isNaN(partsSellingPrice)) {
         printAtMousePointer(this, "输入的配件售价为非法值，请检查后重新提交！");
         $('input[name="partsSellingPriceInput"]').focus();
         return;
      }

      var dataStr = {
         "id": 0,
         "vehicle_name": vehicleName,
         "parts_name": partsName,
         "parts_no": partsNo,
         "parts_count": partsCount,
         "parts_price": partsPrice,
         "parts_selling_price": partsSellingPrice
      };

      var printCtx = ``;
      $.ajax({
         url: '/api/v1/admin/add',
         type: 'post',
         async: false,
         contentType: 'application/json',
         data: JSON.stringify(dataStr),
         success: function (data) {
            if (1 != data.response) {
               // $('.tips').html(`<span style="color:red;">增加失败！, ` + data.msg + `</span>`);
               printCtx = `增加失败！, ` + data.msg;
            } else {
               //在表格最后新增加一行，并填充值
               var row = `<tr align="center">
                                <td style="display:none">` + data.msg.new_id + `</td>
                                <td onclick="tdClick(this)">` + vehicleName + `</td>
                                <td onclick="tdClick(this)">` + partsName + `</td>
                                <td onclick="tdClick(this)">` + partsNo + `</td>
                                <td onclick="tdClick(this)">` + partsCount + `</td>
                                <td onclick="tdClick(this)">` + partsPrice + `</td>
                                <td><input type="button" class="modifyButton" value="修改" onclick="modifyClick(this)"/>&nbsp;&nbsp;&nbsp;&nbsp;
                                <input type="button" class="modifyButton" value="删除" onclick="deleteClick(this)"/></td>
                               </tr>`;
               addRow(row);

               printCtx = `增加成功！`;
               $('input[name="vehicleNameInput"]').val("");
               $('input[name="partsNameInput"]').val("");
               $('input[name="partsNoInput"]').val("");
               $('input[name="partsCountInput"]').val("");
               $('input[name="partsPriceInput"]').val("");
               $('input[name="partsSellingPriceInput"]').val("");

               var oldCount = Number($('#rowCount').text());
               $('#rowCount').text(oldCount + 1);
            }
         }
      });

      printAtMousePointer(this, printCtx);
   });

   //搜索框获得焦点后值变化时 IE
   $('.searchInput').bind("input propertychange", function () {
      var keyWord = $('input[name="searchInput"]').val();
      if ('' == keyWord) {
         doPaging(true, currPageIndex);
      } else {
         adminSearch(keyWord);
      }
   });

   //搜索框获得焦点后值变化时 非IE
   // $('.searchInput').bind("input oninput", function () {
   //     var keyWord = $('input[name="searchInput"]').val();
   //     if ('' == keyWord) {
   //         adminAllparts();
   //     } else {
   //         adminSearch(keyWord);
   //     }
   // });
});

function modifyClick(ele) {
   var id = Number($(ele).parent().siblings('td').eq(0).text());
   var vehicleName = $(ele).parent().siblings('td').eq(1).text();
   var partsName = $(ele).parent().siblings('td').eq(2).text();
   var partsNo = $(ele).parent().siblings('td').eq(3).text();
   var partsCount = Number($(ele).parent().siblings('td').eq(4).text());
   var partsPrice = Number($(ele).parent().siblings('td').eq(5).text());
   var partsSellingPrice = Number($(ele).parent().siblings('td').eq(6).text());

   if (isNaN(id)) {
      printAtMousePointer(ele, "获取ID失败！");
      return;
   }

   if ('' == vehicleName ||
      '' == partsName ||
      '' == partsNo) {
      printAtMousePointer(ele, "输入的必填字段有空值，请检查后重新提交！");
      return;
   }

   if (isNaN(partsCount)) {
      printAtMousePointer(ele, "输入的配件个数为非法值，请检查后重新提交！");
      return;
   }

   if (isNaN(partsPrice)) {
      printAtMousePointer(ele, "输入的配件单价为非法值，请检查后重新提交！");
      return;
   }

   if (isNaN(partsSellingPrice)) {
      printAtMousePointer(ele, "输入的配件售价为非法值，请检查后重新提交！");
      return;
   }

   var dataStr = {
      "id": id,
      "vehicle_name": vehicleName,
      "parts_name": partsName,
      "parts_no": partsNo,
      "parts_count": partsCount,
      "parts_price": partsPrice,
      "parts_selling_price": partsSellingPrice
   };

   var printCtx = ``;
   // 新增的一行
   if (0 == id) {
      $.ajax({
         url: '/api/v1/admin/add',
         type: 'post',
         async: false,
         contentType: 'application/json',
         data: JSON.stringify(dataStr),
         success: function (data) {
            if (1 != data.response) {
               printCtx = `保存失败！` + data.msg;
            } else {
               var newId = data.msg.new_id;
               $(ele).parent().siblings('td').eq(0).text(newId);
               printCtx = `保存成功！`;

               var oldCount = Number($('#rowCount').text());
               $('#rowCount').text(oldCount + 1);

               $(ele).val("修改");
            }
         }
      });
   } else {
      // 一行的修改
      $.ajax({
         url: '/api/v1/admin/update',
         type: 'post',
         async: false,
         contentType: 'application/json',
         data: JSON.stringify(dataStr),
         success: function (data) {
            if (1 != data.response) {
               printCtx = `修改失败！` + data.msg;
            } else {
               printCtx = `修改成功！`;
            }
         }
      });
   }

   $(ele).parent().parent().css("background-color", "#323542");
   $(ele).parent().parent().css("font-size", "12px");
   printAtMousePointer(ele, printCtx);
}

function deleteClick(ele) {
   var id = Number($(ele).parent().siblings('td').eq(0).text());

   if (isNaN(id)) {
      printAtMousePointer(this, "系统异常，请刷新后重试！");
      return;
   }

   // 新增之后未保持，直接删除时
   if (0 == id) {
      $(ele).parent().parent().remove();
      printAtMousePointer(ele, "删除成功！");
      return;
   }

   var dataStr = {
      "id": id,
      "vehicle_name": "",
      "parts_name": "",
      "parts_no": "",
      "parts_count": 0,
      "parts_price": 0.0,
      "parts_selling_price": 0.0
   };

   var printCtx = ``;
   $.ajax({
      url: '/api/v1/admin/delete',
      type: 'post',
      async: false,
      contentType: 'application/json',
      data: JSON.stringify(dataStr),
      success: function (data) {
         if (1 != data.response) {
            printCtx = `删除失败！, ` + data.msg;
         } else {
            $(ele).parent().parent().remove();
            printCtx = `删除成功！`;

            //table取到所有的行
            var rows = $("#resultTable tr").length;
            if (2 == rows) {
               var tb = `<p style="height: 25px;"><span id="rowCount" style="color:#FFF;">共<span style="color:#0BE093;">0</span>条记录</span></p>
                    <table class="resultTable" id="resultTable" frame="border" rules="all">
                         <tr align="center">
                            <th style="display:none">编号</th>
                            <th>车型名称</th>
                            <th>配件名称</th>
                            <th>配件编号</th>
                            <th>数量</th>
                            <th>单价(元)</th>
                            <th>售价(元)</th>
                            <th>操作</th>
                        </tr>
                        <tr>
                        <td colspan="8">
                            <input type="button" class="modifyButton" value="添加" onclick="addNewRow(this)"/></td>
                        </tr>
                    </table>`;
               $(".resultTableDiv").html(tb);
            } else {
               var oldCount = Number($('#rowCount').text());
               $('#rowCount').text(oldCount - 1);
            }
         }
      }
   });

   printAtMousePointer(ele, printCtx);
}

function addNewRow(ele) {
   var row = `<tr align="center" style="background-color: #0BE093; font-size: 20px;">
            <td style="display:none">0</td>
            <td onclick="tdClick(this)"></td>
            <td onclick="tdClick(this)"></td>
            <td onclick="tdClick(this)"></td>
            <td onclick="tdClick(this)"></td>
            <td onclick="tdClick(this)"></td>
            <td onclick="tdClick(this)"></td>
            <td><input type="button" class="modifyButton" value="保存" onclick="modifyClick(this)"/>&nbsp;&nbsp;&nbsp;&nbsp;
                                <input type="button" class="modifyButton" value="删除" onclick="deleteClick(this)"/></td>
        </tr>`;
   addRow(row);
}

function addRow(rowHtml) {
   //table取到所有的行
   var rows = $("#resultTable tr").length;
   var tr = $("#resultTable tr").eq(rows - 2);

   tr.after(rowHtml);
}

function tdClick(ele) {
   if (!$(ele).is('.input')) {
      var oldVal = $(ele).text();
      $(ele).addClass('input').html(`<input type="text" value="` + oldVal + `" style="border: none; background-color: #0BE093; color: inherit; width: 230px; height: 30px; font-size: 20px;"/>`).find('input').focus().blur(function () {
         var afterVal;
         if (1 == $(ele).index() || 2 == $(ele).index() || 3 == $(ele).index()) {
            afterVal = $(this).val();
            if ("" == afterVal) {
               // printAtPos(ele, "修改后的值非数值，请检查后再试！");
               $(this).focus().val(oldVal);
               return;
            }
         }

         // 这里在失去焦点时，对数量和单价的输入做合法性判断
         if (4 == $(ele).index()) {
            afterVal = Number($(this).val());

            // 非数字 非正整数
            if (isNaN(afterVal) || !(/(^[1-9]\d*$)/.test(afterVal))) {
               // printAtPos(ele, "修改后的值非数值，请检查后再试！");
               $(this).focus().val(oldVal);
               return;
            }
         }

         if (5 == $(ele).index()) {
            afterVal = Number($(this).val());
            if (isNaN(afterVal) || 0 > afterVal) {
               // printAtPos(ele, "修改后的值非数值，请检查后再试！");
               $(this).focus().val(oldVal);
               return;
            }
         }

         if (6 == $(ele).index()) {
            afterVal = Number($(this).val());
            if (isNaN(afterVal) || 0 > afterVal) {
               // printAtPos(ele, "修改后的值非数值，请检查后再试！");
               $(this).focus().val(oldVal);
               return;
            }
         }

         if (oldVal != afterVal) {
            $(this).parent().parent().css("background-color", "#0BE093");
         }

         $(this).parent().removeClass('input').html($(this).val() || 0);
      });
   }
}

function adminSearch(keyWord) {
   var dataStr = {
      "key_word": keyWord
   };

   $.ajax({
      url: '/api/v1/admin/search',
      type: 'post',
      async: true,
      contentType: 'application/json',
      data: JSON.stringify(dataStr),
      success: function (data) {
         genTable(data, keyWord, true);
      }
   });
}
