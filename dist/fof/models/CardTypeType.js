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
exports.CardTypeTypeToJSON = exports.CardTypeTypeFromJSONTyped = exports.CardTypeTypeFromJSON = exports.CardTypeType = void 0;
/**
 * This is required for Credit Card Payment Methods. This indicates the type of Credit Card associated with this payment method.
 * @export
 */
exports.CardTypeType = {
    Ab: 'Ab',
    Am: 'Am',
    Ax: 'Ax',
    Cb: 'Cb',
    Dc: 'Dc',
    Ds: 'Ds',
    Dt: 'Dt',
    Ec: 'Ec',
    Er: 'Er',
    Jc: 'Jc',
    Jl: 'Jl',
    Mc: 'Mc',
    Nb: 'Nb',
    So: 'So',
    St: 'St',
    Sw: 'Sw',
    Va: 'Va',
    Xy: 'Xy',
    Zz: 'Zz',
    Cp: 'Cp',
    Cu: 'Cu'
};
function CardTypeTypeFromJSON(json) {
    return CardTypeTypeFromJSONTyped(json, false);
}
exports.CardTypeTypeFromJSON = CardTypeTypeFromJSON;
function CardTypeTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.CardTypeTypeFromJSONTyped = CardTypeTypeFromJSONTyped;
function CardTypeTypeToJSON(value) {
    return value;
}
exports.CardTypeTypeToJSON = CardTypeTypeToJSON;
