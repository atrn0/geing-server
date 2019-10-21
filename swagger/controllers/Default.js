'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.questionsGET = function questionsGET (req, res, next) {
  var page = req.swagger.params['page'].value;
  Default.questionsGET(page)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.questionsPOST = function questionsPOST (req, res, next) {
  var question = req.swagger.params['question'].value;
  Default.questionsPOST(question)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
