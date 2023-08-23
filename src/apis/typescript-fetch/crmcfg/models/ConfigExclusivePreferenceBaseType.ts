/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';

/**
 * Base details common between both template and property level Exclusive preference ids.
 * @export
 * @interface ConfigExclusivePreferenceBaseType
 */
export interface ConfigExclusivePreferenceBaseType {
    /**
     * Specifies the Exclusive preference code.
     * @type {string}
     * @memberof ConfigExclusivePreferenceBaseType
     */
    code?: string;
    /**
     * Specifies the preference group the Exclusive preference belongs to.
     * @type {string}
     * @memberof ConfigExclusivePreferenceBaseType
     */
    preferenceGroup?: string;
    /**
     * Specifies the preference code and its description mapped to the exclusive preference.
     * @type {Array<CodeDescriptionType>}
     * @memberof ConfigExclusivePreferenceBaseType
     */
    preferenceCodes?: Array<CodeDescriptionType>;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ConfigExclusivePreferenceBaseType
     */
    orderSequence?: number;
}

/**
 * Check if a given object implements the ConfigExclusivePreferenceBaseType interface.
 */
export function instanceOfConfigExclusivePreferenceBaseType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ConfigExclusivePreferenceBaseTypeFromJSON(json: any): ConfigExclusivePreferenceBaseType {
    return ConfigExclusivePreferenceBaseTypeFromJSONTyped(json, false);
}

export function ConfigExclusivePreferenceBaseTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigExclusivePreferenceBaseType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'preferenceGroup': !exists(json, 'preferenceGroup') ? undefined : json['preferenceGroup'],
        'preferenceCodes': !exists(json, 'preferenceCodes') ? undefined : ((json['preferenceCodes'] as Array<any>).map(CodeDescriptionTypeFromJSON)),
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
    };
}

export function ConfigExclusivePreferenceBaseTypeToJSON(value?: ConfigExclusivePreferenceBaseType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'preferenceGroup': value.preferenceGroup,
        'preferenceCodes': value.preferenceCodes === undefined ? undefined : ((value.preferenceCodes as Array<any>).map(CodeDescriptionTypeToJSON)),
        'orderSequence': value.orderSequence,
    };
}

