$(document).ready(function () {
  doPaging(false, 1);

  //点击搜索按钮时
  $(".searchButton").click(function (event) {
    var keyWord = $('input[name="searchInput"]').val();
    if ("" == keyWord) {
      return;
    }

    search(keyWord);
  });

  //搜索框获得焦点后值变化时 IE
  $(".searchInput").bind("input propertychange", function () {
    var keyWord = $('input[name="searchInput"]').val();
    if ("" == keyWord) {
      doPaging(false, currPageIndex);
    } else {
      search(keyWord);
    }
  });

  //搜索框获得焦点后值变化时 非IE
  // $('.searchInput').bind("input oninput", function () {
  //     var keyWord = $('input[name="searchInput"]').val();
  //     if ('' == keyWord) {
  //         allparts();
  //     } else {
  //         search(keyWord);
  //     }
  // });
});

function search(keyWord) {
  var dataStr = {
    key_word: keyWord,
  };

  $.ajax({
    url: "/api/v1/index/search",
    type: "post",
    async: true,
    contentType: "application/json",
    data: JSON.stringify(dataStr),
    success: function (data) {
      genTable(data, keyWord, false);
    },
  });
}

function allparts() {
  $.ajax({
    url: "/api/v1/index/allparts",
    type: "post",
    async: true,
    contentType: "application/json",
    data: null,
    success: function (data) {
      genTable(data, "", false);
    },
  });
}
