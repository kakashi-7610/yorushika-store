"use strict";
// alert("test");の場合画面がロードされる前に実行される
// window.addEventListener('load', function () {
//   alert("test");
// });の場合ロード後に実行される

// 削除ボタンが押された時のイベント
const deleteRecommend = () => {
  return confirm("本当に削除しますか？");
}

// const jsonStr = JSON.stringify({
//   name: "テスト",
//   age: 25,
//   interest: ["プログラミング", "料理", "読書"]
// });
// console.log(jsonStr);
// const obj = JSON.parse(jsonStr);
// console.log(obj.name)

// async function displayMessage() {
//   const response = await fetch(testJson);
//   const data = await response.json();
//   const messageElm = document.getElementById('message');
//   messageElm.innerHTML = data.message;
// }

async function getPrefs() {
  const prefResponse = await fetch(testJson);
  return await prefResponse.json();
}

async function displayPrefs() {
  const result = await getPrefs();
  console.log(result);
}

window.addEventListener('load', displayPrefs)