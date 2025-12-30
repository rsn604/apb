# apb について　

　先に作成した**Golang**による**TUI**フレームワーク <a href="https://github.com/rsn604/taps" target="_blank">taps</a> の利用例として、アポイント管理アプリケーション **apb** を作成しました。
いわゆるスケジュール帳で、データは日付と時刻により管理されています。一回限りのスケジュールはもちろんですが、**Day、Week、Month、Year** で繰り返される事項などの登録も簡単にできます。

　画面構成、遷移などは、かつての名機 **HP100LX、200LX**の「**内蔵アプリケーション**」である **Appointment Book** を踏襲しています。

---
## [1] コンパイル

　下記のコマンドでコンパイルします。
```
go build -o apb main.go
```
なお、**bin** ディレクトリに、**Linux、Windows** 用の各バイナリを格納してあります。

---
## [2] 使用法

```
apb <DB Name>
```
で起動します。**DB Name** が存在しない場合は、新規作成します。

---
### (1)　一覧　Apptlist
　当日の予定一覧が表示されます。右上に当月のカレンダー、その下に次の予定と、**TodoList** が同時に表示されています。

![apptlist](/images/01apptlist.png)

なお、以下の主要画面には「**HELP**」が付いているので、**F1**キーやマウスで選択、確認することが出来ます。以降の説明では、キーはファンクションキーのみで説明しますが、その他については「HELP」を参照してください。

![apptlist](/images/01apptlist_help.png)

---
### (2)　データ登録　Detail
　一覧画面から適当な時間を選択すると、「**データ登録**」の画面になります。

![Detail](/images/02detail.png)

必要項目を入力し、**F2** を押します。完了すると、下部に「**D.MG Record added.**」表示されるはずです。

![Detail](/images/02detail_add.png)

---
### (3)　Repeat設定
　レコードの追加が完了すると、「**Repeat設定**」が可能になります。画面下部のガイダンス部に**F8** キーが表示されているはずです。
これを押すと、下記の画面になるので、ここからアポイントの「Repeat設定」を行っていきます。

![Repeat](/images/03repeat_menu.png)

#### 1. Daily
![Repeat](/images/03repeat_daily.png)

　画面を見ていただければわかると思います。**Duration** は、Defaultで5年になっているので、適当な期間を設定してください。

#### 2. Weekly
![Repeat](/images/03repeat_weekly.png)

　「Weekly設定」では、曜日の設定を行います。

#### 3. Monthly
![Repeat](/images/03repeat_monthly.png)

　「Monthly設定」では、日付、あるいは曜日の設定を行います。

#### 4. Yearly
![Repeat](/images/03repeat_yearly.png)

　「Monthly設定」では、日付、あるいは曜日と月の設定を行います。

#### 5. Custom
![Repeat](/images/03repeat_custom.png)

　「Custom設定」は、さらに多くの選択項目があります。

---
### (4)　表示

#### 1. Weekly
　一覧表示から **F8** を押すと週間表示になります。

![View](/images/04wlist.png)

#### 2. Monthly
　一覧表示から **F7** を押すと月間表示になります。

![View](/images/04mlist.png)

---
### (5)　日付の変更

#### 1. Godate
　一覧表示から **F8** を押すと月カレンダーが表示されます。

![Calendar](/images/05godate.png)

この画面から日付を選択、あるいは **Goto** とあるフィールドに直接日付を入力すると、その日付に切り替わります。

#### 2. Calendar(6ヶ月)
　一覧表示から **F7** を押すと6ヶ月カレンダーが表示されます。

![Calendar](/images/05calendar.png)

この画面から日付を選択すると、その日付に切り替わります。

---
### (6)　削除
　一覧項目で、**DEL**キーを押すと、その項目を削除できます。
「Repeat設定」がされていた場合、下記のように3つから選択します。単一スケジュールの場合は、上の2つは選択できない画面となります。

![Delete](/images/06delete.png)

---
### (7)　Todo
　一覧項目から、**F10** を押すと、**TodoList** 一覧が表示されます。新規に追加する場合は、最終のブランク業を選択します。

![Todo](/images/07todolist.png)

登録画面が表示されるので、必要な項目を入力します。**Appointment** 同様、「Repeat設定」も可能です。

![Todo](/images/07todo.png)

---
### (8)　Utility
　**bin** ディレクトリ内に、**Linux、Windows** 用の各バイナリを格納してあります。

#### 1. Unload、Load
　アポイントDB の **Unload** と **Load** が可能です。

```
apbunload <DB Name>
```
データを **JSON**形式で表示します。適当にリダイレクトしてください。

```
apbload <load file> <DB Name>
```
JSON形式のデータをLoadします。なお、既存のDBを指定した場合、同一のデータでも二重に登録してしまいます。あくまで新規DB登録用と考えてください。

#### 2. adbdump、hpload
　HP100LX、200LXからデータを抜き出し、**apb** にコンバートするプログラムです。ただ、実機を持っていないので十分なテストは出来ていないことをご了承ください。参考程度ということでお願いします。

　HPのデータベースからの抜き出しは、<a href="http://mizj.com/" target="_blank">S.U.P.E.R Site</a>というところに登録されている **ADBIO** というツールを使わせてもらいました。

>ADBIO is freeware and may be distributed/copied/used/modified/eaten/microwaved etc. freely. 

とあるので、今回は **gcc** でコンパイルできるように少し変更しています。

まずは、HPのデータをCSVに落とします。**appt.adb** というファイル名は各自の環境によって違うかもしれませんが、HP内のデータを意味します。

```
./adbdump2 -c -i atnr -d "yyyy/mm/dd" -q1 appt.adb >t1.txt
```
ここで、2バイトコードを含んでいる場合は、**nkf** などを使用して、UTF-8の文字列に変換してください。

次に **apb** にロードします。

```
./hpload t1.txt HPL.boltdb
```
あとは、下記のように内容を確認してください。

```
apb HPL.boltdb
```



