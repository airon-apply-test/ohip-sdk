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
import type { AttendantType } from './AttendantType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { EmployeeInfoType } from './EmployeeInfoType';
import type { UniqueIDType } from './UniqueIDType';
import type { UserSessionInfoType } from './UserSessionInfoType';
/**
 *
 * @export
 * @interface ApplicationUserType
 */
export interface ApplicationUserType {
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof ApplicationUserType
     */
    hotel?: CodeDescriptionType;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    appUser?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    lDAPUser?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    actAs?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    actAt?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    userDefaultLanguage?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ApplicationUserType
     */
    appUserId?: UniqueIDType;
    /**
     *
     * @type {number}
     * @memberof ApplicationUserType
     */
    cashierId?: number;
    /**
     * Cashier title.
     * @type {string}
     * @memberof ApplicationUserType
     */
    cashierTitle?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    department?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    departmentLocation?: string;
    /**
     *
     * @type {string}
     * @memberof ApplicationUserType
     */
    salesRepCode?: string;
    /**
     *
     * @type {Date}
     * @memberof ApplicationUserType
     */
    expiryDate?: Date;
    /**
     *
     * @type {Date}
     * @memberof ApplicationUserType
     */
    disabledUntil?: Date;
    /**
     *
     * @type {Date}
     * @memberof ApplicationUserType
     */
    passwordChangeDate?: Date;
    /**
     *
     * @type {EmployeeInfoType}
     * @memberof ApplicationUserType
     */
    userInfo?: EmployeeInfoType;
    /**
     *
     * @type {UserSessionInfoType}
     * @memberof ApplicationUserType
     */
    userSessionInfo?: UserSessionInfoType;
    /**
     *
     * @type {AttendantType}
     * @memberof ApplicationUserType
     */
    attendantInfo?: AttendantType;
}
/**
 * Check if a given object implements the ApplicationUserType interface.
 */
export declare function instanceOfApplicationUserType(value: object): boolean;
export declare function ApplicationUserTypeFromJSON(json: any): ApplicationUserType;
export declare function ApplicationUserTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApplicationUserType;
export declare function ApplicationUserTypeToJSON(value?: ApplicationUserType | null): any;
