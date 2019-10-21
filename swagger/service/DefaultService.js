'use strict';


/**
 * 回答された質問一覧を表示
 * 20件ずつ返す．pageを指定しない場合は最初の20件
 *
 * page BigDecimal ページネーション (optional)
 * returns Questions
 **/
exports.questionsGET = function(page) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = "";
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * 質問を投稿
 *
 * question CreateQuestionRequest 新しい質問
 * returns Question
 **/
exports.questionsPOST = function(question) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "created_at" : "created_at",
  "id" : 0.80082819046101150206595775671303272247314453125,
  "body" : "body"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

