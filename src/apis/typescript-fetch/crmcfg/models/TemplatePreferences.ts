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
import type { ConfigTemplatePreferenceType } from './ConfigTemplatePreferenceType';
import {
    ConfigTemplatePreferenceTypeFromJSON,
    ConfigTemplatePreferenceTypeFromJSONTyped,
    ConfigTemplatePreferenceTypeToJSON,
} from './ConfigTemplatePreferenceType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response object for fetching preferences at the template level.
 * @export
 * @interface TemplatePreferences
 */
export interface TemplatePreferences {
    /**
     * This type holds a collection of preferences at the template level.
     * @type {Array<ConfigTemplatePreferenceType>}
     * @memberof TemplatePreferences
     */
    templatePreferences?: Array<ConfigTemplatePreferenceType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof TemplatePreferences
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof TemplatePreferences
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the TemplatePreferences interface.
 */
export function instanceOfTemplatePreferences(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TemplatePreferencesFromJSON(json: any): TemplatePreferences {
    return TemplatePreferencesFromJSONTyped(json, false);
}

export function TemplatePreferencesFromJSONTyped(json: any, ignoreDiscriminator: boolean): TemplatePreferences {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'templatePreferences': !exists(json, 'templatePreferences') ? undefined : ((json['templatePreferences'] as Array<any>).map(ConfigTemplatePreferenceTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function TemplatePreferencesToJSON(value?: TemplatePreferences | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'templatePreferences': value.templatePreferences === undefined ? undefined : ((value.templatePreferences as Array<any>).map(ConfigTemplatePreferenceTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

