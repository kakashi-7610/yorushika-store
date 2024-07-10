"use strict";
// alert("test");の場合画面がロードされる前に実行される
// window.addEventListener('load', function () {
//   alert("test");
// });の場合ロード後に実行される


class MovieViewer {
  // 設計の柔軟性を考え、画像最大数やどの要素にビューアーを作るかは
  // 可変にしたいため、画像と要素は引数で渡すようにする
  constructor(rootElm, movies) {
    this.rootElm = rootElm;
    this.movies = movies;
    this.currentIndex = 0;
  }

  init() {
    // const nextButtonElm = this.rootElm.querySelector('.nextButton');
    // nextButtonElm.addEventListener('click', () => {
    //   this.next()
    // })

    // const prevButtonElm = this.rootElm.querySelector('.prevButton');
    // prevButtonElm.addEventListener('click', () => {
    //   this.prev()
    // })

    this.updateMovie();
  }

  updateMovie() {
    const movie = this.movies[this.currentIndex];
    this.rootElm.innerHTML = movie;

    this.startTimer();
  }

  startTimer() {
    // すでにタイマーが動作している場合は停止しないと、インターバル処理が
    // 複数個実行されてしまう。
    if (this.timerKey) {
      clearTimeout(this.timerKey)
    }

    // timeoutIDをローカル変数で持ってしまうと関数が呼び出される度にリセットされてしまう。
    // インスタンス変数で持つことでインスタンスが解放されるまで利用できる。

    // setTimeoutは時間切れになったら１度だけ関数を実行する
    // setIntervalは設定した時間間隔で永遠に関数が実行される

    // setIntervalでの繰り返し処理
    // 処理にかかる時間に関係なく指定した秒数ごとに関数が実行される。

    // setTimeoutでの繰り返し処理
    // 関数が終了してから指定した秒数後に関数が実行される。
    this.timerKey = setTimeout(() => {
      this.next();
    }, 5000);
  }

  next() {
    this.currentIndex++;
    if (this.currentIndex === this.movies.length) {
      this.currentIndex = 0;
    }

    this.updateMovie();
  }

  prev() {
    this.currentIndex--;
    if (this.currentIndex < 0) {
      this.currentIndex = this.movies.length - 1;
    }

    this.updateMovie();
  }
}


const execution = () => {
  const rootElm = document.querySelector(".movie-frame");
  const movies = [
    '<iframe class="movie" src="https://www.youtube.com/embed/Ht6lcYg9Zfo" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>',
    '<iframe class="movie src="//www.youtube.com/embed/F64yFFnZfkI?rel=0" frameborder="0" allowfullscreen="allowfullscreen"></iframe>',
    '<iframe class="movie" src="https://www.youtube.com/embed/Sw1Flgub9s8" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>'
  ]

  const moviewViewer = new MovieViewer(rootElm, movies);
  moviewViewer.init();
}

window.addEventListener('load', execution)