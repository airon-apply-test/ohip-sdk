/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { EmailDeliveryFormatType } from './EmailDeliveryFormatType';
import type { EmailDeliveryMethodType } from './EmailDeliveryMethodType';
/**
 * Email content configuration template
 * @export
 * @interface EmailDeliveryConfigurationType
 */
export interface EmailDeliveryConfigurationType {
    /**
     *
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    fromAddress?: string;
    /**
     *
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    userId?: string;
    /**
     *
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    userPassword?: string;
    /**
     *
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    serverName?: string;
    /**
     *
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    subject?: string;
    /**
     *
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    body?: string;
    /**
     *
     * @type {EmailDeliveryFormatType}
     * @memberof EmailDeliveryConfigurationType
     */
    format?: EmailDeliveryFormatType;
    /**
     *
     * @type {EmailDeliveryMethodType}
     * @memberof EmailDeliveryConfigurationType
     */
    type?: EmailDeliveryMethodType;
    /**
     * Hotel code
     * @type {string}
     * @memberof EmailDeliveryConfigurationType
     */
    hotelId?: string;
    /**
     * This option must be selected in order to activate emailing from an SMTP compliant email server
     * @type {boolean}
     * @memberof EmailDeliveryConfigurationType
     */
    activateDelivery?: boolean;
    /**
     * Available for Confirmation Letters only and when Activate Email Delivery is selected
     * @type {boolean}
     * @memberof EmailDeliveryConfigurationType
     */
    hTMLFormatDelivery?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof EmailDeliveryConfigurationType
     */
    attachICalender?: boolean;
}
/**
 * Check if a given object implements the EmailDeliveryConfigurationType interface.
 */
export declare function instanceOfEmailDeliveryConfigurationType(value: object): boolean;
export declare function EmailDeliveryConfigurationTypeFromJSON(json: any): EmailDeliveryConfigurationType;
export declare function EmailDeliveryConfigurationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailDeliveryConfigurationType;
export declare function EmailDeliveryConfigurationTypeToJSON(value?: EmailDeliveryConfigurationType | null): any;
