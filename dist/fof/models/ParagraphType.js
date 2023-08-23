"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ParagraphTypeToJSON = exports.ParagraphTypeFromJSONTyped = exports.ParagraphTypeFromJSON = exports.instanceOfParagraphType = void 0;
const runtime_1 = require("../runtime");
const FormattedTextTextType_1 = require("./FormattedTextTextType");
/**
 * Check if a given object implements the ParagraphType interface.
 */
function instanceOfParagraphType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfParagraphType = instanceOfParagraphType;
function ParagraphTypeFromJSON(json) {
    return ParagraphTypeFromJSONTyped(json, false);
}
exports.ParagraphTypeFromJSON = ParagraphTypeFromJSON;
function ParagraphTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'text': !(0, runtime_1.exists)(json, 'text') ? undefined : (0, FormattedTextTextType_1.FormattedTextTextTypeFromJSON)(json['text']),
        'image': !(0, runtime_1.exists)(json, 'image') ? undefined : json['image'],
        'url': !(0, runtime_1.exists)(json, 'url') ? undefined : json['url'],
    };
}
exports.ParagraphTypeFromJSONTyped = ParagraphTypeFromJSONTyped;
function ParagraphTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'text': (0, FormattedTextTextType_1.FormattedTextTextTypeToJSON)(value.text),
        'image': value.image,
        'url': value.url,
    };
}
exports.ParagraphTypeToJSON = ParagraphTypeToJSON;
