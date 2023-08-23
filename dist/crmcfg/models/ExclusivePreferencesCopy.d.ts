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
import type { ConfigCopyExclusivePreferencesType } from './ConfigCopyExclusivePreferencesType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for copying template Exclusive preferences to hotel(s).
 * @export
 * @interface ExclusivePreferencesCopy
 */
export interface ExclusivePreferencesCopy {
    /**
     *
     * @type {ConfigCopyExclusivePreferencesType}
     * @memberof ExclusivePreferencesCopy
     */
    copyInstructions?: ConfigCopyExclusivePreferencesType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ExclusivePreferencesCopy
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ExclusivePreferencesCopy
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ExclusivePreferencesCopy interface.
 */
export declare function instanceOfExclusivePreferencesCopy(value: object): boolean;
export declare function ExclusivePreferencesCopyFromJSON(json: any): ExclusivePreferencesCopy;
export declare function ExclusivePreferencesCopyFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExclusivePreferencesCopy;
export declare function ExclusivePreferencesCopyToJSON(value?: ExclusivePreferencesCopy | null): any;
