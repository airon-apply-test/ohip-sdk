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
exports.PostingRhythmTypeToJSON = exports.PostingRhythmTypeFromJSONTyped = exports.PostingRhythmTypeFromJSON = exports.PostingRhythmType = void 0;
/**
 * Simple type for posting Rhythm types.
 * @export
 */
exports.PostingRhythmType = {
    EveryNight: 'EveryNight',
    ArrivalNight: 'ArrivalNight',
    EveryXNightsStartingNightY: 'EveryXNightsStartingNightY',
    CertainNightsOfTheWeek: 'CertainNightsOfTheWeek',
    LastNight: 'LastNight',
    EveryNightExceptArrivalNight: 'EveryNightExceptArrivalNight',
    EveryNightExceptLast: 'EveryNightExceptLast',
    EveryNightExceptFirstAndLast: 'EveryNightExceptFirstAndLast',
    CustomStaySchedule: 'CustomStaySchedule',
    CustomNightSchedule: 'CustomNightSchedule',
    FloatingAllowancePerStay: 'FloatingAllowancePerStay',
    TicketPosting: 'TicketPosting'
};
function PostingRhythmTypeFromJSON(json) {
    return PostingRhythmTypeFromJSONTyped(json, false);
}
exports.PostingRhythmTypeFromJSON = PostingRhythmTypeFromJSON;
function PostingRhythmTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.PostingRhythmTypeFromJSONTyped = PostingRhythmTypeFromJSONTyped;
function PostingRhythmTypeToJSON(value) {
    return value;
}
exports.PostingRhythmTypeToJSON = PostingRhythmTypeToJSON;
