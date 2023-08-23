/**
 * Opera Cloud Inventory Asynchronous API
 * APIs to cater for Inventory related asynchronous functionality in OPERA Cloud. This includes viewing inventory data along with its revenue and updating inventory&apos;s sell limits for date ranges. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.4.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface RoomClassSellLimitType
 */
export interface RoomClassSellLimitType {
    /**
     *
     * @type {Date}
     * @memberof RoomClassSellLimitType
     */
    date?: Date;
    /**
     *
     * @type {number}
     * @memberof RoomClassSellLimitType
     */
    amount?: number;
    /**
     * Indicates if sell limit is flat or percentage.
     * @type {string}
     * @memberof RoomClassSellLimitType
     */
    flatOrPercentage?: RoomClassSellLimitTypeFlatOrPercentageEnum;
    /**
     *
     * @type {string}
     * @memberof RoomClassSellLimitType
     */
    roomClass?: string;
}
/**
 * @export
 */
export declare const RoomClassSellLimitTypeFlatOrPercentageEnum: {
    readonly Flat: "Flat";
    readonly Percentage: "Percentage";
};
export type RoomClassSellLimitTypeFlatOrPercentageEnum = typeof RoomClassSellLimitTypeFlatOrPercentageEnum[keyof typeof RoomClassSellLimitTypeFlatOrPercentageEnum];
/**
 * Check if a given object implements the RoomClassSellLimitType interface.
 */
export declare function instanceOfRoomClassSellLimitType(value: object): boolean;
export declare function RoomClassSellLimitTypeFromJSON(json: any): RoomClassSellLimitType;
export declare function RoomClassSellLimitTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomClassSellLimitType;
export declare function RoomClassSellLimitTypeToJSON(value?: RoomClassSellLimitType | null): any;
