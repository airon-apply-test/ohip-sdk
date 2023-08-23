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
/**
 *
 * @export
 * @interface HSKTaskSheetInfoTypeRoomSummaryInner
 */
export interface HSKTaskSheetInfoTypeRoomSummaryInner {
    /**
     * Simple type for room status instructions to be used in requests for fetching housekeeping rooms. Valid status values are Clean, Dirty, Pickup, Inspected, OutOfOrder, OutOfService.
     * @type {string}
     * @memberof HSKTaskSheetInfoTypeRoomSummaryInner
     */
    housekeepingRoomStatusType?: HSKTaskSheetInfoTypeRoomSummaryInnerHousekeepingRoomStatusTypeEnum;
    /**
     *
     * @type {number}
     * @memberof HSKTaskSheetInfoTypeRoomSummaryInner
     */
    totalCount?: number;
}
/**
 * @export
 */
export declare const HSKTaskSheetInfoTypeRoomSummaryInnerHousekeepingRoomStatusTypeEnum: {
    readonly Clean: "Clean";
    readonly Dirty: "Dirty";
    readonly Pickup: "Pickup";
    readonly Inspected: "Inspected";
    readonly OutOfOrder: "OutOfOrder";
    readonly OutOfService: "OutOfService";
};
export type HSKTaskSheetInfoTypeRoomSummaryInnerHousekeepingRoomStatusTypeEnum = typeof HSKTaskSheetInfoTypeRoomSummaryInnerHousekeepingRoomStatusTypeEnum[keyof typeof HSKTaskSheetInfoTypeRoomSummaryInnerHousekeepingRoomStatusTypeEnum];
/**
 * Check if a given object implements the HSKTaskSheetInfoTypeRoomSummaryInner interface.
 */
export declare function instanceOfHSKTaskSheetInfoTypeRoomSummaryInner(value: object): boolean;
export declare function HSKTaskSheetInfoTypeRoomSummaryInnerFromJSON(json: any): HSKTaskSheetInfoTypeRoomSummaryInner;
export declare function HSKTaskSheetInfoTypeRoomSummaryInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): HSKTaskSheetInfoTypeRoomSummaryInner;
export declare function HSKTaskSheetInfoTypeRoomSummaryInnerToJSON(value?: HSKTaskSheetInfoTypeRoomSummaryInner | null): any;
