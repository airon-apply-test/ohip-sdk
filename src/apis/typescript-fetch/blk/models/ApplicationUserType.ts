/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AttendantType } from './AttendantType';
import {
    AttendantTypeFromJSON,
    AttendantTypeFromJSONTyped,
    AttendantTypeToJSON,
} from './AttendantType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { EmployeeInfoType } from './EmployeeInfoType';
import {
    EmployeeInfoTypeFromJSON,
    EmployeeInfoTypeFromJSONTyped,
    EmployeeInfoTypeToJSON,
} from './EmployeeInfoType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';
import type { UserSessionInfoType } from './UserSessionInfoType';
import {
    UserSessionInfoTypeFromJSON,
    UserSessionInfoTypeFromJSONTyped,
    UserSessionInfoTypeToJSON,
} from './UserSessionInfoType';

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
export function instanceOfApplicationUserType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ApplicationUserTypeFromJSON(json: any): ApplicationUserType {
    return ApplicationUserTypeFromJSONTyped(json, false);
}

export function ApplicationUserTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApplicationUserType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotel': !exists(json, 'hotel') ? undefined : CodeDescriptionTypeFromJSON(json['hotel']),
        'appUser': !exists(json, 'appUser') ? undefined : json['appUser'],
        'lDAPUser': !exists(json, 'lDAPUser') ? undefined : json['lDAPUser'],
        'actAs': !exists(json, 'actAs') ? undefined : json['actAs'],
        'actAt': !exists(json, 'actAt') ? undefined : json['actAt'],
        'userDefaultLanguage': !exists(json, 'userDefaultLanguage') ? undefined : json['userDefaultLanguage'],
        'appUserId': !exists(json, 'appUserId') ? undefined : UniqueIDTypeFromJSON(json['appUserId']),
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
        'cashierTitle': !exists(json, 'cashierTitle') ? undefined : json['cashierTitle'],
        'department': !exists(json, 'department') ? undefined : json['department'],
        'departmentLocation': !exists(json, 'departmentLocation') ? undefined : json['departmentLocation'],
        'salesRepCode': !exists(json, 'salesRepCode') ? undefined : json['salesRepCode'],
        'expiryDate': !exists(json, 'expiryDate') ? undefined : (new Date(json['expiryDate'])),
        'disabledUntil': !exists(json, 'disabledUntil') ? undefined : (new Date(json['disabledUntil'])),
        'passwordChangeDate': !exists(json, 'passwordChangeDate') ? undefined : (new Date(json['passwordChangeDate'])),
        'userInfo': !exists(json, 'userInfo') ? undefined : EmployeeInfoTypeFromJSON(json['userInfo']),
        'userSessionInfo': !exists(json, 'userSessionInfo') ? undefined : UserSessionInfoTypeFromJSON(json['userSessionInfo']),
        'attendantInfo': !exists(json, 'attendantInfo') ? undefined : AttendantTypeFromJSON(json['attendantInfo']),
    };
}

export function ApplicationUserTypeToJSON(value?: ApplicationUserType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotel': CodeDescriptionTypeToJSON(value.hotel),
        'appUser': value.appUser,
        'lDAPUser': value.lDAPUser,
        'actAs': value.actAs,
        'actAt': value.actAt,
        'userDefaultLanguage': value.userDefaultLanguage,
        'appUserId': UniqueIDTypeToJSON(value.appUserId),
        'cashierId': value.cashierId,
        'cashierTitle': value.cashierTitle,
        'department': value.department,
        'departmentLocation': value.departmentLocation,
        'salesRepCode': value.salesRepCode,
        'expiryDate': value.expiryDate === undefined ? undefined : (value.expiryDate.toISOString().substr(0,10)),
        'disabledUntil': value.disabledUntil === undefined ? undefined : (value.disabledUntil.toISOString().substr(0,10)),
        'passwordChangeDate': value.passwordChangeDate === undefined ? undefined : (value.passwordChangeDate.toISOString().substr(0,10)),
        'userInfo': EmployeeInfoTypeToJSON(value.userInfo),
        'userSessionInfo': UserSessionInfoTypeToJSON(value.userSessionInfo),
        'attendantInfo': AttendantTypeToJSON(value.attendantInfo),
    };
}

