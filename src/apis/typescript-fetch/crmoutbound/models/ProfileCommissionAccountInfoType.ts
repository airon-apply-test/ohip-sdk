/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BankAccountType } from './BankAccountType';
import {
    BankAccountTypeFromJSON,
    BankAccountTypeFromJSONTyped,
    BankAccountTypeToJSON,
} from './BankAccountType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { ProfileId } from './ProfileId';
import {
    ProfileIdFromJSON,
    ProfileIdFromJSONTyped,
    ProfileIdToJSON,
} from './ProfileId';

/**
 * Profile commission info which contains bank account and commission code details
 * @export
 * @interface ProfileCommissionAccountInfoType
 */
export interface ProfileCommissionAccountInfoType {
    /**
     * 
     * @type {ProfileId}
     * @memberof ProfileCommissionAccountInfoType
     */
    profileId?: ProfileId;
    /**
     * 
     * @type {BankAccountType}
     * @memberof ProfileCommissionAccountInfoType
     */
    bankAccount?: BankAccountType;
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof ProfileCommissionAccountInfoType
     */
    commissionCode?: CodeDescriptionType;
}

/**
 * Check if a given object implements the ProfileCommissionAccountInfoType interface.
 */
export function instanceOfProfileCommissionAccountInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileCommissionAccountInfoTypeFromJSON(json: any): ProfileCommissionAccountInfoType {
    return ProfileCommissionAccountInfoTypeFromJSONTyped(json, false);
}

export function ProfileCommissionAccountInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileCommissionAccountInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileId': !exists(json, 'profileId') ? undefined : ProfileIdFromJSON(json['profileId']),
        'bankAccount': !exists(json, 'bankAccount') ? undefined : BankAccountTypeFromJSON(json['bankAccount']),
        'commissionCode': !exists(json, 'commissionCode') ? undefined : CodeDescriptionTypeFromJSON(json['commissionCode']),
    };
}

export function ProfileCommissionAccountInfoTypeToJSON(value?: ProfileCommissionAccountInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileId': ProfileIdToJSON(value.profileId),
        'bankAccount': BankAccountTypeToJSON(value.bankAccount),
        'commissionCode': CodeDescriptionTypeToJSON(value.commissionCode),
    };
}

