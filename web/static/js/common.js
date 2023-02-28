var timeout = 4000;     // 操作提示消失的毫秒数
var pagingNum = 10;     // 多少行为一页
var currPageIndex = 1;  // 当前页号

function logoutClick() {
    $.ajax({
        url: '/api/v1/account/logout',
        type: 'post',
        async: true,
        contentType: 'application/json',
        data: null,
        success: function (data) {
            if (1 != data.response) {
                alert(`退出失败！，` + data.msg);
                return;
            } else {
                location.href = data.msg;
            }
        }
    });
}

function printAtMousePointer(ele, text) {
    var event = $(ele).event || window.event;
    var x = event.pageX || event.clientX + document.body.scroolLeft;
    var y = event.pageY || event.clientY + document.body.scrollTop;
    var ctx = $("<span />").text(text);
    ctx.css({
        "z-index": 999999999999999999999999999999999999999999999999999999999999999999999,
        "top": y - 20,
        "left": x,
        "position": "absolute",
        "font-weight": "bold",
        "color": "#ff6651",
        "font-size": "16px"
    });
    $("body").append(ctx);
    ctx.animate({
            "top": y - 90,
            "opacity": 0
        },
        timeout,
        function () {
            ctx.remove();
        });
    return;
}

function printAtPos(ele, text) {
    var x = ele.offset().top;
    var y = ele.offset().left;
    var ctx = $("<span />").text(text);
    ctx.css({
        "z-index": 999999999999999999999999999999999999999999999999999999999999999999999,
        "top": y - 20,
        "left": x,
        "position": "absolute",
        "font-weight": "bold",
        "color": "#ff6651",
        "font-size": "16px"
    });
    $("body").append(ctx);
    ctx.animate({
            "top": y - 90,
            "opacity": 0
        },
        timeout,
        function () {
            ctx.remove();
        });
    return;
}

function doPaging(isAll, index) {
    var dataStr = {
        "limit_paging_num": pagingNum,
        "index": index
    }

    var url = "";
    if (isAll) {
        url = `/api/v1/admin/paging`;
    } else {
        url = `/api/v1/index/paging`;
    }

    $.ajax({
        url: url,
        type: 'post',
        async: true,
        contentType: 'application/json',
        data: JSON.stringify(dataStr),
        success: function (data) {
            genTable(data, "", isAll);

            // 修改页号颜色
            var oldPageIndex = currPageIndex;
            currPageIndex = index;
            $("#pageIndexId" + oldPageIndex).css("color", "#FFF");
            $("#pageIndexId" + currPageIndex).css("color", "#0BE093");
        }
    });

}

function adminAllparts() {
    $.ajax({
        url: "/api/v1/admin/allparts",
        type: 'post',
        async: true,
        contentType: 'application/json',
        data: null,
        success: function (data) {
            genTable(data, "", isAll);
        }
    });
}

function genTable(tData, keyWord, isAll) {
    if (1 != tData.response) {
        if ('' != keyWord) {
            $(".resultTableDiv").html(`<h2>无[` + keyWord + `]对应的搜索结果，请尝试其它关键字</h2>`);
            return;
        } else {
            $(".resultTableDiv").html(`<h2>获取配件信息失败!</h2>`);
            return;
        }
    }

    var rowCount = tData.msg.row_count;
    if (0 == rowCount) {
        if ('' != keyWord) {
            $(".resultTableDiv").html(`<h2>无[` + keyWord + `]对应的搜索结果，请尝试其它关键字</h2>`);
            return;
        } else {
            if (isAll) {
                var tb = `<p style="height: 25px;"><span style="color:#FFF;">共<span id="rowCount" style="color:#0BE093;">0</span>条记录</span></p>
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
                        <td colspan="7">
                            <input type="button" class="modifyButton" value="添加" onclick="addNewRow(this)"/></td>
                        </tr>
                    </table>`;
                $(".resultTableDiv").html(tb);
            } else {
                $(".resultTableDiv").html(`<h2>数据库中无数据！等待录入中……</h2>`);
            }
            return;
        }
    }

    var tableStr = "";
    if ("" == keyWord) {
        // 分页，非搜索
        var totalCount = 0;
        $.ajax({
            url: "/api/v1/index/count",
            type: 'post',
            async: false,
            contentType: 'application/json',
            data: null,
            success: function (data) {
                if (1 != data.response) {
                    alert(`分页失败！，` + data.msg);
                    return;
                } else {
                    totalCount = data.msg.count;
                }
            }
        });

        var pagingTotal = Math.ceil(totalCount / pagingNum); // 总分页数

        tableStr += `<p style="height: 25px; font-size: 14px; padding-top: 8px;">
                <span style="color:#FFF; padding-top: 8px;">共<span id="rowCount" style="color:#0BE093;">` + totalCount + `</span>行</span>
                <span style="color:#FFF;">分<span id="rowCount" style="color:#0BE093;">` + pagingTotal + `</span>页</span>
                <a style="color:#FFF;" href="#" onclick="doPaging(` + isAll + `, 1)">首页</a>&nbsp;&nbsp;`;

        for (var i = 1; i <= pagingTotal; i++) {
            tableStr += `<a id="pageIndexId` + i + `" style="color:#FFF;" href="#" onclick="doPaging(` + isAll + `, ` + i + `)">` + i + `</a>&nbsp;&nbsp;`;
        }

        tableStr += `<a style="color:#FFF;" href="#" onclick="doPaging(` + isAll + `, ` + pagingTotal + `)">尾页</a>&nbsp;&nbsp;</p>`;
        // 分页结束
    } else {
        // 搜索时不分页
        tableStr += `<p style="height: 25px; font-size: 14px; padding-top: 8px;">
                <span style="color:#FFF; padding-top: 8px;">共<span id="rowCount" style="color:#0BE093;">` + rowCount + `</span>行</span>&nbsp;&nbsp;</p>`;
    }

    tableStr += `<table class="resultTable" id="resultTable" frame="border" rules="all">`;

    if (isAll) {
        tableStr += `<tr align="center">
                <th style="display:none">编号</th>
                <th>车型名称</th>
                <th>配件名称</th>
                <th>配件编号</th>
                <th>数量</th>
                <th>单价(元)</th>
                <th>售价(元)</th>
                <th>操作</th>
            </tr>`;

        for (var i = 0; i < rowCount; i++) {
            tableStr += `<tr align="center">
                <td style="display:none">` + tData.msg.row_infos[i].id + `</td>
                <td onclick="tdClick(this)">` + tData.msg.row_infos[i].vehicle_name + `</td>
                <td onclick="tdClick(this)">` + tData.msg.row_infos[i].parts_name + `</td>
                <td onclick="tdClick(this)">` + tData.msg.row_infos[i].parts_no + `</td>
                <td onclick="tdClick(this)">` + tData.msg.row_infos[i].parts_count + `</td>
                <td onclick="tdClick(this)">` + tData.msg.row_infos[i].parts_price + `</td>
                <td onclick="tdClick(this)">` + tData.msg.row_infos[i].parts_selling_price + `</td>
                <!--<td onclick="tdClick(this)"><input value="7" style="border: none; background-color: inherit; color: inherit;"></td>-->
                <td><input type="button" class="modifyButton" value="修改" onclick="modifyClick(this)"/>&nbsp;&nbsp;&nbsp;&nbsp;
                <input type="button" class="modifyButton" value="删除" onclick="deleteClick(this)"/></td>
                </tr>`;
        }

        tableStr += `<tr>
            <td colspan="7">
            <input type="button" class="modifyButton" value="添加" onclick="addNewRow(this)"/></td>
        </tr>`;
    } else {
        tableStr += `<tr align="center">
                <th style="display:none">编号</th>
                <th>车型名称</th>
                <th>配件名称</th>
                <th>配件编号</th>
                <th>数量</th>
                <th>售价(元)</th>
                </tr>`;

        for (var i = 0; i < rowCount; i++) {
            tableStr += `<tr align="center">
                <td style="display:none">` + tData.msg.row_infos[i].id + `</td>
                <td>` + tData.msg.row_infos[i].vehicle_name + `</td>
                <td>` + tData.msg.row_infos[i].parts_name + `</td>
                <td>` + tData.msg.row_infos[i].parts_no + ` </td>
                <td>` + tData.msg.row_infos[i].parts_count + `</td>
                <td>` + tData.msg.row_infos[i].parts_selling_price + `</td>
                </tr>`;
        }
    }

    tableStr += `</table>`;
    $(".resultTableDiv").html(tableStr);
}