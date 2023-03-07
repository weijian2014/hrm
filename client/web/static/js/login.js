$(function () {
  //点击
  $(".loginBtn").click(function (event) {
    login();
  });

  function login() {
    var username = $('input[name="username"]').val();
    if ("" == username || "用户名" == username) {
      alert("请输入用户名！");
      return;
    }

    var pwd = $('input[name="password"]').val();
    if (pwd == "") {
      alert("请输入密码！");
      return;
    }

    $.ajax({
      url: "/api/v1/account/login",
      dataType: "json",
      type: "post",
      data: { username: username, password: pwd },
      success: function (loginData) {
        var loginStarus = loginData.response;
        var loginUrl = loginData.msg;
        if (loginStarus == 1) {
          location.href = loginUrl;
        } else {
          alert(loginUrl);
        }
      },
    });
  }
});
