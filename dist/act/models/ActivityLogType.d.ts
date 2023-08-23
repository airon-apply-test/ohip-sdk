/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ActivityLogType
 */
export interface ActivityLogType {
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    hotelId?: string;
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    module?: string;
    /**
     *
     * @type {Date}
     * @memberof ActivityLogType
     */
    logDate?: Date;
    /**
     *
     * @type {number}
     * @memberof ActivityLogType
     */
    refActionId?: number;
    /**
     *
     * @type {number}
     * @memberof ActivityLogType
     */
    logUserId?: number;
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    logUserName?: string;
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    machineStation?: string;
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    croCode?: string;
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    actionType?: string;
    /**
     *
     * @type {string}
     * @memberof ActivityLogType
     */
    actionDescription?: string;
    /**
     * The IP Address of the machine that performed the activity
     * @type {string}
     * @memberof ActivityLogType
     */
    iPAddress?: string;
}
/**
 * Check if a given object implements the ActivityLogType interface.
 */
export declare function instanceOfActivityLogType(value: object): boolean;
export declare function ActivityLogTypeFromJSON(json: any): ActivityLogType;
export declare function ActivityLogTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityLogType;
export declare function ActivityLogTypeToJSON(value?: ActivityLogType | null): any;
