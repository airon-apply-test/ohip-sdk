"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ArticlePostItTypeToJSON = exports.ArticlePostItTypeFromJSONTyped = exports.ArticlePostItTypeFromJSON = exports.instanceOfArticlePostItType = void 0;
const runtime_1 = require("../runtime");
const ColorType_1 = require("./ColorType");
/**
 * Check if a given object implements the ArticlePostItType interface.
 */
function instanceOfArticlePostItType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfArticlePostItType = instanceOfArticlePostItType;
function ArticlePostItTypeFromJSON(json) {
    return ArticlePostItTypeFromJSONTyped(json, false);
}
exports.ArticlePostItTypeFromJSON = ArticlePostItTypeFromJSON;
function ArticlePostItTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'availableForPostIt': !(0, runtime_1.exists)(json, 'availableForPostIt') ? undefined : json['availableForPostIt'],
        'color': !(0, runtime_1.exists)(json, 'color') ? undefined : (0, ColorType_1.ColorTypeFromJSON)(json['color']),
    };
}
exports.ArticlePostItTypeFromJSONTyped = ArticlePostItTypeFromJSONTyped;
function ArticlePostItTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'availableForPostIt': value.availableForPostIt,
        'color': (0, ColorType_1.ColorTypeToJSON)(value.color),
    };
}
exports.ArticlePostItTypeToJSON = ArticlePostItTypeToJSON;
